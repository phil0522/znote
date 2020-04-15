package subcommands

import (
	"flag"
	"strings"

	"github.com/phil0522/znote/pkg/notesmarket"
	pb "github.com/phil0522/znote/proto"
	"github.com/sirupsen/logrus"
)

var (
	EditNoteCommandFlagSet = flag.NewFlagSet("Edit", flag.ExitOnError)
)

func CreateEditNoteRequest() pb.ZNoteRequest {
	if EditNoteCommandFlagSet.NArg() == 0 {
		logrus.Panic("need note id")
	}

	return pb.ZNoteRequest{
		Command: "edit",
		NoteId:  EditNoteCommandFlagSet.Arg(0),
	}
}

func ResolveEditNote(req pb.ZNoteRequest) pb.ZNoteResponse {
	noteId := req.NoteId

	market := notesmarket.GetNotesMarket()

	sb := &strings.Builder{}
	for _, book := range market.Books {
		if book.Notes.HasNoteById(noteId) {
			go func() {
				note := book.Notes.GetNote(noteId)
				book.EditNote(note)
				market.SaveAll()
			}()
		}
	}

	return pb.ZNoteResponse{
		Result: sb.String(),
	}
}
