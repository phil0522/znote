package notesmarket

import (
	"path/filepath"
	"time"

	"github.com/phil0522/znote/pkg/config"
	"github.com/sirupsen/logrus"
)

func (b *Book) getFilePathFromCreationTime(timeValue string) string {
	currentTime, err := time.Parse(time.RFC3339, timeValue)
	if err != nil {
		logrus.WithField("time", timeValue).Fatal("failed to parse creation time")
		return ""
	}
	return filepath.Join(config.RootDir, b.Name, "ByMonth", currentTime.Format("200601.md"))
}

func (b *Book) getFilePathForProject(project string) string {
	return filepath.Join(config.RootDir, b.Name, "ByProject", project)
}

func (b *Book) getFilePathForTag(tag string) string {
	return filepath.Join(config.RootDir, b.Name, "ByTags", tag)
}
