package main

import (
	"flag"
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
list: list all notes
pick: pick a note and edit it
edit: edit a note directly
`
)

type subCommand struct {
	name    string
	flagset *flag.FlagSet
	method  func()
}

func makeSubCommand(name string, flagset *flag.FlagSet, method func()) subCommand {
	return subCommand{
		name:    name,
		flagset: flagset,
		method:  method,
	}
}

var (
	commands = []subCommand{
		makeSubCommand("new", subcommands.NewNoteCommandFlagSet, subcommands.NewNote),
		makeSubCommand("list", subcommands.ListNoteCommandFlagSet, subcommands.ListNote),
		makeSubCommand("edit", subcommands.EditNoteCommandFlagSet, subcommands.EditNote),
		makeSubCommand("pick", subcommands.PickNoteCommandFlagSet, subcommands.PickNoteForEdit),
		makeSubCommand("serve", subcommands.ServeCommandFlagSet, subcommands.ServeHttp),
		makeSubCommand("refresh", subcommands.RefreshCommandFlagSet, subcommands.RefreshNotes),
	}
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	logrus.WithField("args", os.Args[2:]).Info("params")
	for _, command := range commands {
		if command.name == os.Args[1] {
			err := command.flagset.Parse(os.Args[2:])
			if err != nil {
				logrus.WithError(err).Panic("parse flag failure.")
			}
			command.method()
			return
		}
	}
	fmt.Printf("%s is not a valid command\n", os.Args[1])
}
