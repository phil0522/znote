package notesmarket

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Book may contains multiple notes. Each book has its own directory.
// Notes may be duplicated to different locations.
// ByProject, ByTags, ByMonth
type Book struct {
	Name  string
	Notes NoteSet

	noteFiles map[string]*NoteFile
}

func NewBook() *Book {
	return &Book{
		Notes:     NewNoteSet(),
		noteFiles: make(map[string]*NoteFile),
	}
}
func (b *Book) RemoveNote(n *Note) {
	notePath := b.getFilePathFromCreationTime(n.CreationTime)
	nf := b.noteFiles[notePath]
	nf.notes.removeNoteByKey(n.Id)
	b.Notes.removeNoteByKey(n.Id)
}

func (b *Book) EditNote(n Note) bool {
	updated := editNote(&n)
	if updated == nil {
		logrus.Debug("no change")
		return false
	}
	b.RemoveNote(&n)
	if updated.Id == IdToRemove {
		return true
	}

	if updated.Project != "" && updated.Project != b.Name {
		otherBook := GetNotesMarket().GetOrCreateBook(updated.Project)
		otherBook.AddNote(*updated)
		otherBook.saveToDisk()
	} else {
		b.AddNote(*updated)
	}
	return true
}

func (b *Book) AddNote(n Note) {
	logrus.WithField("book", b.Name).WithField("note", n).Info("Add new Note")
	notePath := b.getFilePathFromCreationTime(n.CreationTime)
	nf, ok := b.noteFiles[notePath]
	if !ok {
		nf = NewNoteFile(notePath)
	}
	nf.notes.upsertNode(n)
	nf.changed = true
	b.noteFiles[notePath] = nf

	b.Notes.upsertNode(n)
}

func (b *Book) saveToDisk() {
	for _, nf := range b.noteFiles {
		nf.save()
	}
}

func (b *Book) UpdateBook() {
	f, err := os.Create(filepath.Join(RootDir, b.Name, b.Name+".toc"))

	if err != nil {
		logrus.WithError(err).Panic("failed to update root.toc")
	}
	defer f.Close()

	fmt.Fprintln(f, "ByMonth,")
	for name := range b.noteFiles {
		relName := b.getRelPath(name)
		fmt.Fprintf(f, "%s,%s\n", relName, titleForByMonthPath(relName))
	}

	fmt.Fprintln(f, "ByTag,")
	for tag := range b.UpdateTags() {
		fmt.Fprintf(f, "ByTag/%s.md,%s\n", tag, tag)
	}
}

func (b *Book) getRelPath(p string) string {
	bookBaseDir := filepath.Join(RootDir, b.Name)
	relPath, err := filepath.Rel(bookBaseDir, p)
	if err != nil {
		logrus.WithField("base", bookBaseDir).WithField("path", p).Panic("can not get relative path")
	}
	return relPath
}

func titleForByMonthPath(path string) string {
	_, fileName := filepath.Split(path)
	baseName := strings.Split(fileName, ".")[0]

	if strings.Contains(path, "ByMonth") {
		yearMonth, err := time.Parse("200601", baseName)
		if err != nil {
			logrus.WithField("path", path).WithField("base", baseName).Panic("can not parse")
		}
		return yearMonth.Format("2006 January")
	}
	return "Not Supported"
}

func (b *Book) UpdateTags() map[string][]*Note {
	noteByTags := make(map[string][]*Note)

	for _, note := range b.Notes.notes {
		if note.Archived {
			continue
		}
		for _, tag := range note.Tags {
			noteByTags[tag] = append(noteByTags[tag], note)
		}
	}

	for tag, notes := range noteByTags {
		noteFile := NewNoteFile(filepath.Join(RootDir, b.Name, "ByTag", strings.ToLower(tag)+".md"))
		noteFile.notes.mergeNotes(notes)
		noteFile.changed = true
		noteFile.save()
	}
	return noteByTags
}
