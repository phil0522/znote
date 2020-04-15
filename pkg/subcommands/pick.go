package subcommands

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

var (
	PickNoteCommandFlagSet = flag.NewFlagSet("Pick", flag.ExitOnError)
)

func PickNoteForEdit() {
	logrus.Debug("List Note")

	bookName, noteId := invokeFzfAndGetBookNoteId()
	if bookName == "" || noteId == "" {
		return
	}

	market := notesmarket.GetNotesMarket()

	book := market.GetOrCreateBook(bookName)
	note := book.Notes.GetNote(noteId)

	if book.EditNote(note) {
		market.SaveAll()
	}
}

func invokeFzfAndGetBookNoteId() (bookName string, noteId string) {
	market := notesmarket.GetNotesMarket()

	cmd := exec.Command("fzf", "--height", "50%", "--border", "-d,", "--preview", "znote show {1} {2}", "--color", "fg:-1,bg:-1,hl:230,fg+:3,bg+:233,hl+:229")

	pipeInput, err := cmd.StdinPipe()
	if err != nil {
		logrus.WithError(err).Panic("failed to get pipe input")
	}
	pipeOutput, err := cmd.StdoutPipe()
	if err != nil {
		logrus.WithError(err).Panic("failed to get pipe output")
	}
	cmd.Stderr = os.Stderr
	for k, v := range market.Books {
		logrus.WithField("book", k).Debug("Book")
		for _, note := range v.Notes.ToOrderedList() {
			fmt.Fprintf(pipeInput, "%s,%s,%s\n", k, note.CreationTime, note.Title)
		}
	}

	if err = cmd.Start(); err != nil {
		logrus.WithError(err).Panic("failed to start fzf ")
	}
	logrus.Debug("started fzf")
	_ = pipeInput.Close()

	logrus.Debug("closed pipeinput")
	bytes, err := ioutil.ReadAll(pipeOutput)
	if err != nil {
		logrus.WithError(err).Panic("failed to read output")
	}
	text := strings.TrimSpace(string(bytes))
	if text == "" {
		logrus.Debug("doesn't pick, do nothing")
		return "", ""
	}
	logrus.WithField("output", text).Info("get fzf output")

	fields := strings.Split(text, ",")
	return fields[0], fields[1]
}
