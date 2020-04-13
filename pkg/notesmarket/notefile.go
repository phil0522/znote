package notesmarket

import (
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

// noteFile keep several notes in one markdown file
type noteFile struct {
	content string
	notes   []Note
}

func (n *noteFile) load(notePath string) {
	bytes, err := ioutil.ReadFile(notePath)
	if err != nil {
		logrus.Fatalf("failed to read file %s", notePath)
	}
	n.content = string(bytes)
	n.notes = make([]Note, 0)

	// Ignore the line with single #, which is the title of the file
	// Start with first ##
	lines := strings.Split(n.content, "\n")

	note := Note{}
	for _, line := range lines {
		// logrus.WithField("line", line).Info("Get line")
		if strings.HasPrefix(line, "## ") {
			if note.Title != "" {
				n.notes = append(n.notes, note)
			}
			note = Note{}
			note.Title = strings.Trim(string(line[3:]), "\n")
			continue
		}
		note.Content = note.Content + line + "\n"
	}
	if note.Title != "" {
		n.notes = append(n.notes, note)
	}
	logrus.WithField("notes", n.notes).Info("get notes")
}

func (n *noteFile) save(notePath string) {

}
