package subcommands

import (
	"fmt"
	"math/rand"
	"os/exec"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/sirupsen/logrus"
)

func maybeSubmit() {
	cmdStr := fmt.Sprintf("cd %s && if [[ `git status --porcelain` ]]; then git add . ; git commit -a -m auto-submit; fi", notesmarket.RootDir)

	cmd := exec.Command("/bin/zsh", "-c", cmdStr)

	if bytes, err := cmd.CombinedOutput(); err != nil {
		logrus.WithError(err).Warnf("Failed to commit\n%s\n", string(bytes))
	}
	maybePush()
	logrus.Debug("submit changes")
}

func maybePush() {
	r := rand.Intn(10000000)
	if r != 75 {
		return
	}
	cmdStr := fmt.Sprintf("cd %s && git push", notesmarket.RootDir)

	cmd := exec.Command("/bin/zsh", "-c", cmdStr)

	if bytes, err := cmd.CombinedOutput(); err != nil {
		logrus.WithError(err).Warnf("Failed to push\n%s\n", string(bytes))
	}
	logrus.Debug("push commits to server")
}
