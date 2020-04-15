package subcommands

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

var (
	RefreshCommandFlagSet = flag.NewFlagSet("Refresh", flag.ExitOnError)
)

func RefreshNotes() {
	logrus.Debug("Refresh notes to get latest state.")
	updateToc()
}

func updateToc() {
	logrus.Debug("update toc")
	market := notesmarket.GetNotesMarket()

	f, err := os.Create(filepath.Join(notesmarket.RootDir, "root.toc"))

	if err != nil {
		logrus.WithError(err).Panic("failed to update root.toc")
	}
	defer f.Close()

	for name, book := range market.Books {
		fmt.Fprintln(f, name)
		book.UpdateBookToc()
	}
}
