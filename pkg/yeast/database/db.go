package database

import (
	"path/filepath"

	"github.com/phil0522/znote/pkg/config"
	"github.com/sirupsen/logrus"
)

type Database struct {
}

const (
	databaseRelPath = "database/notes.db"
)

func init() {
	dbPath := filepath.Join(config.RootDir, databaseRelPath)
	logrus.WithField("dbPath", dbPath).Info("Open database")

	db, err := bolt.Open(dbPath, 0600, nil)

	if err != nil {
		logrus.WithError(err).WithField("dbPath", dbPath).Warn("Failed to open database")
	}
}
