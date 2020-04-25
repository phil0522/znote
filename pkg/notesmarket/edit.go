package notesmarket

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/phil0522/znote/config"
	"github.com/sirupsen/logrus"
)

// Creates a temporary file and open external editor to edit, load the edit content
func editNote(note *Note) *Note {
	if note.Id == "" {
		logrus.Panic("note id can not be empty")
	}
	tmpFilePath := filepath.Join(config.RootDir, "tmp", note.Id+".md")
	nf := NewNoteFile(tmpFilePath)

	nf.notes.upsertNode(*note)
	nf.changed = true
	nf.save()

	err := openFileInEditor(nf.notePath)
	if err != nil {
		logrus.WithField("path", nf.notePath).Warning("Error edit notes")
		return nil
	}
	nf.load()
	_ = os.Remove(nf.notePath)
	if len(nf.notes.notes) != 1 {
		logrus.Info("invalid input format, discard change")
		return nil
	}
	return nf.notes.ToOrderedList()[0]
}

const (
	DefaultEditor = "vim"
)

func openFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	if strings.Contains(executable, "Visual Studio Code") || strings.Contains(executable, "vscode") {
		cmd = exec.Command(executable, "--wait", filename)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
