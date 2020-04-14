package subcommands

import (
	"flag"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

// NewNoteCommandFlagSet is flag set for creating note.
var NewNoteCommandFlagSet = flag.NewFlagSet("NewNote", flag.ExitOnError)

func NewNote() {
	bookName := NewNoteCommandFlagSet.Arg(0)
	logrus.WithField("book", bookName).Info("Create Note")
	market := notesmarket.GetNotesMarket()
	book := market.GetOrCreateBook(bookName)

	var n notesmarket.Note
	if NewNoteCommandFlagSet.NArg() == 1 {
		n = notesmarket.EmptyNote()
		n.Title = "Test 1"
	} else {
		n = newNoteWithTemplate(NewNoteCommandFlagSet.Arg(1))
	}
	book.AddNote(n)
	book.EditNote(n)
	market.SaveAll()
}

func newNoteWithTemplate(templateName string) notesmarket.Note {
	baseNote := notesmarket.EmptyNote()
	switch templateName {
	case "task":
		baseNote.Title = ":active: Title"
		baseNote.Tags = append(baseNote.Tags, "task-active")
	default:
		logrus.WithField("template", templateName).Panic("unknown template name")
	}
	return baseNote
}
