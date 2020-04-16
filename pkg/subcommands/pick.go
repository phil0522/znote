package subcommands

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	pb "github.com/phil0522/znote/proto"
	"github.com/sirupsen/logrus"
)

var (
	PickNoteCommandFlagSet = flag.NewFlagSet("Pick", flag.ExitOnError)
)

func CreatePickNoteRequest() pb.ZNoteRequest {
	noteId := invokeFzfAndGetBookNoteId()

	return pb.ZNoteRequest{
		Command: "edit",
		NoteId:  noteId,
	}
}

func ResolvePickNote(req pb.ZNoteRequest) pb.ZNoteResponse {
	panic("Not Implemented")
}

func invokeFzfAndGetBookNoteId() string {

	cmd2Input, cmd1Output := io.Pipe()

	cmd1 := exec.Command("znote", "list")
	//cmd2 := exec.Command("fzf")
	cmd2 := exec.Command("fzf", "--height", "50%", "--border", "-d,", "--preview", "znote show {2}", "--color", "fg:-1,bg:-1,hl:230,fg+:3,bg+:233,hl+:229")

	cmd1.Stdout = cmd1Output
	cmd2.Stdin = cmd2Input
	cmd2.Stderr = os.Stderr
	pipeOutput, err := cmd2.StdoutPipe()
	if err != nil {
		logrus.WithError(err).Panic("failed to open fzf output")
	}

	if err := cmd1.Start(); err != nil {
		logrus.WithError(err).Panic("failed to run znote")
	}

	if err := cmd2.Start(); err != nil {
		logrus.WithError(err).Panic("failed to start fzf ")
	}

	if err := cmd1.Wait(); err != nil {
		logrus.WithError(err).Panic("failed to wait cmd1")
	}
	_ = cmd1Output.Close()

	go func() {
		if err := cmd2.Wait(); err != nil {
			logrus.WithError(err).Warn("failed to wait cmd2")
		}
	}()

	logrus.Info("read output")
	bytes, err := ioutil.ReadAll(pipeOutput)
	if err != nil {
		logrus.WithError(err).Panic("failed to read output")
	}
	logrus.Info("done")
	text := strings.TrimSpace(string(bytes))
	if text == "" {
		logrus.Debug("doesn't pick, do nothing")
		return ""
	}
	logrus.WithField("output", text).Info("get fzf output")

	fields := strings.Split(text, ",")
	return fields[1]
}
