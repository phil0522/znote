package subcommands

import (
	"flag"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

var (
	EditNoteCommandFlagSet = flag.NewFlagSet("Edit", flag.ExitOnError)
)

func EditNote() {
	logrus.Info("List Note")
	market := notesmarket.GetNotesMarket()

	book := market.GetOrCreateBook(EditNoteCommandFlagSet.Arg(0))

	noteId := EditNoteCommandFlagSet.Arg(1)

	note := book.Notes.GetNote(noteId)
	book.EditNote(note)

	market.SaveAll()
}
