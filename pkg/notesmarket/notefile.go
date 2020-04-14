package notesmarket

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// NoteFile keep several notes in one markdown file
type NoteFile struct {
	notePath string
	content  string
	notes    NoteSet
	changed  bool
}

func (n *NoteFile) load() {
	bytes, err := ioutil.ReadFile(n.notePath)
	if err != nil {
		logrus.Fatalf("failed to read file %s", n.notePath)
	}
	n.content = string(bytes)
	n.notes = NewNoteSet()

	// Ignore the line with single #, which is the title of the file
	// Start with first ##
	lines := strings.Split(n.content, "\n")

	note := Note{}
	for _, line := range lines {
		// logrus.WithField("line", line).Info("Get line")
		if strings.HasPrefix(line, "## ") {
			if note.Title != "" {
				note.parseContent()
				n.notes.upsertNode(note)
			}
			note = Note{}
			note.Title = strings.Trim(string(line[3:]), "\n")
			continue
		}
		note.Content = note.Content + line + "\n"

	}
	if note.Title != "" {
		note.parseContent()
		n.notes.upsertNode(note)
	}
	logrus.WithField("path", n.notePath).WithField("notes", n.notes).Debug("Loaded with notes")
}

func (nf *NoteFile) save() {
	logger := logrus.WithField("notefile", nf.notePath)
	if !nf.changed {
		logger.Info("No change, skip saving")
		return
	}

	logger.Info("save")

	notes := nf.notes.ToOrderedList()

	baseDir, _ := filepath.Split(nf.notePath)
	_ = os.MkdirAll(baseDir, 0755)
	f, err := os.Create(nf.notePath)
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()
	for _, note := range notes {
		fmt.Fprintln(f, note.contentToSave())
	}
}

func NewNoteFile(notePath string) *NoteFile {
	logrus.WithField("notePath", notePath).Debug("new note file")
	nf := &NoteFile{}
	nf.notes = NewNoteSet()
	nf.notePath = notePath
	return nf
}
