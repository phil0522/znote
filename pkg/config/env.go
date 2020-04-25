package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var (
	RootDir = getRootDir()
)

func getRootDir() string {
	p := os.Getenv("ZNOTES_ROOT")
	if p == "" {
		p = filepath.Join(os.Getenv("HOME"), "znotes", "public")
		logrus.WithField("default_dir", p).Warn("Env ZNOTES_ROOT is not set, use default value")
	}
	logrus.WithField("root", p).Warn("Env ZNOTES_ROOT is not set, use default value")
	return p
}
