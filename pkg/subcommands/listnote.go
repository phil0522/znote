package subcommands

import (
	"flag"
	"fmt"
	"strings"

	"github.com/phil0522/znote/pkg/notesmarket"

	pb "github.com/phil0522/znote/proto"
	"github.com/sirupsen/logrus"
)

var (
	ListNoteCommandFlagSet = flag.NewFlagSet("List", flag.ExitOnError)
)

func CreateListNoteRequest() pb.ZNoteRequest {
	return pb.ZNoteRequest{
		Command: "list",
	}
}
func ResolveListNote(req pb.ZNoteRequest) pb.ZNoteResponse {
	logrus.Debug("List Note")
	market := notesmarket.GetNotesMarket()
	sb := &strings.Builder{}
	for k, v := range market.Books {
		logrus.WithField("book", k).Debug("Book")
		for _, note := range v.Notes.ToOrderedList() {
			fmt.Fprintf(sb, "%s,%s,%s,%s\n", k, note.Id, note.Title, strings.Join(note.Tags, " "))
		}
	}
	resp := pb.ZNoteResponse{}
	resp.Result = sb.String()
	return resp
}
