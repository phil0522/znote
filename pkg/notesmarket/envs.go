package notesmarket

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

var (
	// RootDir is the root dir
	RootDir = getRootDir()
)

func getRootDir() string {
	p := os.Getenv("ZNOTES_ROOT")
	if p == "" {
		p = path.Join(os.Getenv("HOME"), "znotes", "public")
		logrus.Warnf("Env ZNOTES_ROOT is not set, use default value %s", p)
	}
	return p
}
