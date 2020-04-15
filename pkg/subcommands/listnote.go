package subcommands

import (
	"flag"
	"fmt"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

var (
	ListNoteCommandFlagSet = flag.NewFlagSet("List", flag.ExitOnError)
)

func ListNote() {
	logrus.Debug("List Note")
	market := notesmarket.GetNotesMarket()
	for k, v := range market.Books {
		logrus.WithField("book", k).Debug("Book")
		for _, note := range v.Notes.ToOrderedList() {
			fmt.Printf("%s,%s,%s\n", k, note.Id, note.Title)
		}
	}
}
