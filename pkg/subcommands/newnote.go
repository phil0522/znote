package subcommands

import (
	"flag"
)

// NewNoteCommandFlagSet is flag set for creating note.
var NewNoteCommandFlagSet = flag.NewFlagSet("NewNote", flag.ExitOnError)

func NewNote() {
	logger.Warn("List Note")
}
