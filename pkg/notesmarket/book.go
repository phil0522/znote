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
	nf.notes.removeNoteByKey(n.CreationTime)
	b.Notes.removeNoteByKey(n.CreationTime)
}

func (b *Book) EditNote(n Note) bool {
	updated := editNote(&n)
	if updated == nil {
		logrus.Debug("no change")
		return false
	}
	b.RemoveNote(&n)
	b.AddNote(*updated)
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

func (b *Book) UpdateBookToc() {
	f, err := os.Create(filepath.Join(RootDir, b.Name, b.Name+".toc"))

	if err != nil {
		logrus.WithError(err).Panic("failed to update root.toc")
	}
	defer f.Close()

	first := true
	for name := range b.noteFiles {
		relName := b.getRelPath(name)
		if first {
			fmt.Fprintf(f, "%s,\n", getBaseDirectory(relName))
			first = false
		}
		fmt.Fprintf(f, "%s,%s", relName, titleForPath(relName))
	}
}

func getBaseDirectory(p string) string {
	base, _ := filepath.Split(p)
	return strings.TrimRight(base, "/")
}

func (b *Book) getRelPath(p string) string {
	bookBaseDir := filepath.Join(RootDir, b.Name)
	relPath, err := filepath.Rel(bookBaseDir, p)
	if err != nil {
		logrus.WithField("base", bookBaseDir).WithField("path", p).Panic("can not get relative path")
	}
	return relPath
}

func titleForPath(path string) string {
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
