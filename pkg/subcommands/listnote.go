package subcommands

import (
	"flag"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

var (
	logger                 = logrus.WithField("ROOT", notesmarket.RootDir)
	ListNoteCommandFlagSet = flag.NewFlagSet("List", flag.ExitOnError)
)

func ListNote() {
	logrus.Warn("List Note")
	market := notesmarket.GetNotesMarket()
	for k, v := range market.Books {
		logrus.WithField("book", k).Info("Book")
		for _, note := range v.Notes {
			logrus.WithField("title", note.Title).WithField("content", note.Content).Infof("note")
		}
	}
}
