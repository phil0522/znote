package notesmarket

import (
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	nextId int
)

var (
	iniPath string
	cfg     *ini.File
)

func init() {
	iniPath = filepath.Join(RootDir, "znote.ini")

	var err error
	cfg, err = ini.LoadSources(ini.LoadOptions{Loose: true}, iniPath)
	if err != nil {
		logrus.WithField("ini", iniPath).Panic("failed to load file")
	}

	nextId = cfg.Section("notes").Key("nextId").MustInt(100000)
	logrus.WithField("nextId", nextId).Debug("initialize nextId")
}

func generateNextId() string {
	nextId += 1
	cfg.Section("notes").Key("nextId").SetValue(strconv.Itoa(nextId))
	_ = cfg.SaveTo(iniPath)

	return strconv.Itoa(nextId)
}
