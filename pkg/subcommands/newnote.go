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

	n := notesmarket.EmptyNote()
	n.Title = "Test 1"
	book.AddNote(n)

	market.SaveAll()
}
