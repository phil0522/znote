package notesmarket

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var (
	// RootDir is the root dir
	RootDir = getRootDir()
)

func getRootDir() string {
	p := os.Getenv("ZNOTES_ROOT")
	if p == "" {
		p = filepath.Join(os.Getenv("HOME"), "znotes", "public")
		logrus.Warnf("Env ZNOTES_ROOT is not set, use default value %s", p)
	}
	return p
}
