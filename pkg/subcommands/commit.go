package subcommands

import (
	"fmt"
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
}
