package main

import (
	"fmt"
	"os"

	"github.com/phil0522/znote/pkg/subcommands"
	"github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

const (
	usage = `Usage: znote <command> [<args>]

command can be of
new: create a new note
ls: list all notes
`
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	switch os.Args[1] {
	case "new":
		_ = subcommands.NewNoteCommandFlagSet.Parse(os.Args[2:])
	case "list":
		_ = subcommands.ListNoteCommandFlagSet.Parse(os.Args[2:])
	case "edit":
		_ = subcommands.EditNoteCommandFlagSet.Parse(os.Args[2:])
	case "pick":
		_ = subcommands.PickNoteCommandFlagSet.Parse(os.Args[2:])
	default:
		fmt.Printf("%s is not a valid command\n", os.Args[1])
		return
	}

	if subcommands.NewNoteCommandFlagSet.Parsed() {
		subcommands.NewNote()
	} else if subcommands.ListNoteCommandFlagSet.Parsed() {
		subcommands.ListNote()
	} else if subcommands.EditNoteCommandFlagSet.Parsed() {
		subcommands.EditNote()
	} else if subcommands.PickNoteCommandFlagSet.Parsed() {
		subcommands.PickNoteForEdit()
	}
}
