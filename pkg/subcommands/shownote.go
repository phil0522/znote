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

	showNotePreview = ShowNoteCommandFlagSet.Bool("preview", false, "show only first bytes.")
)

func CreateShowNoteRequest() pb.ZNoteRequest {
	if ShowNoteCommandFlagSet.NArg() == 0 {
		logrus.Panic("need note id")
	}

	if *showNotePreview {
		return pb.ZNoteRequest{
			Command: "preview",
			NoteId:  ShowNoteCommandFlagSet.Arg(0),
		}
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
			content := note.Content
			if req.Command == "preview" && len(content) > 300 {
				content = content[0:300]
			}
			fmt.Fprintf(sb, "## %s\n@%s\n%s\n", note.Title, strings.Join(note.Tags, ", "), content)
		}
	}
	return pb.ZNoteResponse{
		Result: sb.String(),
	}
}
