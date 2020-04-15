package notesmarket

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Note is the basic unit of the notes system.
type Note struct {
	Title        string
	Content      string
	Project      string
	Tags         []string
	CreationTime string // RFC3339 format of creation time, used as unique key
	Archived     bool
}

func EmptyNote() Note {
	return Note{
		CreationTime: time.Now().Format(time.RFC3339),
	}
}
func (n *Note) copyFrom(other Note) {
	n.Title = other.Title
	n.Content = other.Content
	n.Project = other.Project
	n.Tags = other.Tags
	n.Archived = other.Archived
}

const (
	znoteLinePrefix = "znote:"
)

func (n *Note) parseContent() {
	realContent := make([]string, 0)
	stage := "title"
	for _, line := range strings.Split(n.Content, "\n") {
		if stage == "title" {
			stage = "header"
			continue
		}
		if stage == "header" {
			if strings.HasPrefix(line, "-->") {
				stage = "body"
				continue
			}
			if strings.HasPrefix(line, "<!--") {
				continue
			}
			n.updateFromPropertiesLine(line)
		}
		if stage == "body" {
			realContent = append(realContent, line)
		}
	}

	n.Content = strings.Join(realContent, "")
}

func (n *Note) updateFromPropertiesLine(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	if !strings.HasPrefix(line, znoteLinePrefix) {
		logrus.WithField("line", line).Warn("failed to parse line")
		panic("not a valid znote property line")
	}
	line = line[len(znoteLinePrefix):]

	fields := strings.SplitN(line, "=", 2)

	logrus.WithField("field", fields).Debug("get Fields")
	key := strings.TrimSpace(fields[0])
	value := strings.TrimSpace(fields[1])

	switch key {
	case "created":
		n.CreationTime = value
	case "tags":
		fields = strings.Split(value, ",")
		for _, field := range fields {
			tag := strings.TrimSpace(field)
			if tag != "" {
				n.Tags = append(n.Tags, tag)
			}
		}
	case "project":
		n.Project = value
	case "status":
		n.Archived = false
		if value == "archived" {
			n.Archived = true
		}
	default:
		logrus.WithField("field", fields).Warn("unknown keys")
		panic("unknown keys")
	}
}

func (n *Note) headerText() string {
	status := "active"
	if n.Archived {
		status = "archived"
	}
	return fmt.Sprintf(`<!--
znote: created=%s
znote: project=%s
znote: tags=%s
znote: status=%s
-->`, n.CreationTime, n.Project, strings.Join(n.Tags, ","), status)
}

func (n *Note) contentToSave() string {
	return fmt.Sprintf("## %s\n%s\n%s\n", n.Title, n.headerText(), n.Content)
}
