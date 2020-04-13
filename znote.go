package main

import (
	"fmt"
	"os"

	"github.com/phil0522/znote/pkg/subcommands"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`
usage: znote <command> [<args>]
		`)
		return
	}

	switch os.Args[1] {
	case "new":
		subcommands.NewNoteCommandFlagSet.Parse(os.Args[2:])
	default:
		fmt.Printf("%s is not a valid command\n", os.Args[1])
		return
	}

	if subcommands.NewNoteCommandFlagSet.Parsed() {
		fmt.Println("New Node")
	}
}
