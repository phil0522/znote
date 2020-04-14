package subcommands

import (
	"flag"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

var (
	ListNoteCommandFlagSet = flag.NewFlagSet("List", flag.ExitOnError)
)

func ListNote() {
	logrus.Info("List Note")
	market := notesmarket.GetNotesMarket()
	for k, v := range market.Books {
		logrus.WithField("book", k).Info("Book")
		for _, note := range v.Notes.ToOrderedList() {
			logrus.WithField("title", note.Title).WithField("CreationTime", note.CreationTime).Infof("note")
		}
	}
}
