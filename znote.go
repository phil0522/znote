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
		subcommands.NewNoteCommandFlagSet.Parse(os.Args[2:])
		break
	case "list":
		subcommands.ListNoteCommandFlagSet.Parse(os.Args[2:])
	default:
		fmt.Printf("%s is not a valid command\n", os.Args[1])
		return
	}

	if subcommands.NewNoteCommandFlagSet.Parsed() {
		logrus.Info("New Note")
	} else if subcommands.ListNoteCommandFlagSet.Parsed() {
		subcommands.ListNote()
	}
}
