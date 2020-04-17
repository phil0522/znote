package subcommands

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/phil0522/znote/pkg/notesmarket"
	pb "github.com/phil0522/znote/proto"
	"github.com/sirupsen/logrus"
)

var (
	RefreshCommandFlagSet = flag.NewFlagSet("Refresh", flag.ExitOnError)
)

func CreateRefreshMarketRequest() pb.ZNoteRequest {
	return pb.ZNoteRequest{
		Command: "refresh",
	}
}
func ResolveRefreshMarket(req pb.ZNoteRequest) pb.ZNoteResponse {
	logrus.Debug("Refresh Market")
	go updateToc()
	return pb.ZNoteResponse{}
}

func updateToc() {
	logrus.Debug("update toc")
	market := notesmarket.GetNotesMarket()
	market.Reload()

	f, err := os.Create(filepath.Join(notesmarket.RootDir, "root.toc"))

	if err != nil {
		logrus.WithError(err).Panic("failed to update root.toc")
	}
	defer f.Close()

	for name, book := range market.Books {
		fmt.Fprintln(f, name)
		book.UpdateBook()
	}

	for name, book := range market.PageBooks {
		fmt.Fprintln(f, name)
		book.ReloadAll()
		book.UpdateToc()
	}

	maybeSubmit()
}
