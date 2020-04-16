package subcommands

import (
	"fmt"
	"math/rand"
	"os/exec"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

func maybeSubmit() {
	cmdStr := fmt.Sprintf("cd %s && git commit -a -m auto-submit", notesmarket.RootDir)

	cmd := exec.Command("/bin/zsh", "-c", cmdStr)

	if err := cmd.Run(); err != nil {
		logrus.WithError(err).Warn("Failed to commit")
	}
	maybePush()
}

func maybePush() {
	r := rand.Intn(100)
	if r != 75 {
		return
	}
	cmdStr := fmt.Sprintf("cd %s && git push", notesmarket.RootDir)

	cmd := exec.Command("/bin/zsh", "-c", cmdStr)

	if err := cmd.Run(); err != nil {
		logrus.WithError(err).Warn("Failed to push")
	}
}
