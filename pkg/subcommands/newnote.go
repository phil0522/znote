package subcommands

import (
	"flag"

	"github.com/phil0522/znote/pkg/notesmarket"
	pb "github.com/phil0522/znote/proto"
	"github.com/sirupsen/logrus"
)

// NewNoteCommandFlagSet is flag set for creating note.
var NewNoteCommandFlagSet = flag.NewFlagSet("NewNote", flag.ExitOnError)

var (
	bookFlag = NewNoteCommandFlagSet.String("b", "work", "the book where note is created.")
	template = NewNoteCommandFlagSet.String("t", "default", "template to use")
)

func NewNoteCreateRequest() pb.ZNoteRequest {
	req := pb.ZNoteRequest{}
	req.Command = "new"
	req.Book = *bookFlag
	return req
}

func ResolveNewNote(request pb.ZNoteRequest) pb.ZNoteResponse {
	bookName := NewNoteCommandFlagSet.Arg(0)
	logrus.WithField("book", bookName).WithField("args", NewNoteCommandFlagSet.Args()).Info("Create Note")
	market := notesmarket.GetNotesMarket()
	book := market.GetOrCreateBook(bookName)

	var n notesmarket.Note = newNoteWithTemplate(*template)
	book.AddNote(n)
	if book.EditNote(n) {
		market.SaveAll()
	} else {
		book.RemoveNote(&n)
	}
	return pb.ZNoteResponse{}
}

func newNoteWithTemplate(templateName string) notesmarket.Note {
	baseNote := notesmarket.EmptyNote()
	switch templateName {
	case "task":
		baseNote.Title = ":active: Title"
		baseNote.Tags = append(baseNote.Tags, "task-active")
	default:
		logrus.WithField("template", templateName).Warn("unknown template name")
	}
	return baseNote
}
