package subcommands

import (
	"flag"
	"fmt"
	"strings"

	"github.com/phil0522/znote/pkg/notesmarket"
)

var (
	ShowNoteCommandFlagSet = flag.NewFlagSet("List", flag.ExitOnError)
)

func ShowNote() {
	bookName := ShowNoteCommandFlagSet.Arg(0)
	noteId := ShowNoteCommandFlagSet.Arg(1)

	market := notesmarket.GetNotesMarket()

	book := market.GetOrCreateBook(bookName)

	if book.Notes.HasNoteByCreationTime(noteId) {
		note := book.Notes.GetNote(noteId)
		fmt.Printf("## %s\n@%s\n%s\n", note.Title, strings.Join(note.Tags, ", "), note.Content)
	}
}
