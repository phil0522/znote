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
	ShowNoteCommandFlagSet = flag.NewFlagSet("List", flag.ExitOnError)
)

func CreateShowNoteRequest() pb.ZNoteRequest {
	if ShowNoteCommandFlagSet.NArg() == 0 {
		logrus.Panic("need note id")
	}

	return pb.ZNoteRequest{
		Command: "show",
		NoteId:  ShowNoteCommandFlagSet.Arg(0),
	}
}

func ResolveShowNote(req pb.ZNoteRequest) pb.ZNoteResponse {
	noteId := req.NoteId

	market := notesmarket.GetNotesMarket()

	sb := &strings.Builder{}
	for _, book := range market.Books {
		if book.Notes.HasNoteById(noteId) {
			note := book.Notes.GetNote(noteId)
			fmt.Fprintf(sb, "## %s\n@%s\n%s\n", note.Title, strings.Join(note.Tags, ", "), note.Content)
		}
	}
	return pb.ZNoteResponse{
		Result: sb.String(),
	}
}
