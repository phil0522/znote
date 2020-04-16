package notesmarket

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Note is the basic unit of the notes system.
type Note struct {
	Id           string
	Title        string
	Content      string
	Project      string
	Tags         []string
	CreationTime string // RFC3339 format of creation time, used as unique key
	Archived     bool
}

func EmptyNote() Note {
	return Note{
		Id:           generateNextId(),
		CreationTime: time.Now().Format(time.RFC3339),
	}
}

const (
	znoteLinePrefix = "znote: "
	IdToRemove      = "removed"
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
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "-->") {
				stage = "body"
				continue
			}
			if line == "" || strings.HasPrefix(line, "<!--") {
				continue
			}
			if strings.HasPrefix(line, znoteLinePrefix) {
				n.updateFromPropertiesLine(line)
			} else {
				stage = "body"
			}
		}
		if stage == "body" {
			if !n.consumeInlineInstruction(line) {
				realContent = append(realContent, line+"\n")
			}
		}
	}

	n.Content = strings.Join(realContent, "")

	if n.Id == "" {
		n.Id = generateNextId()
	}
}

func (n *Note) consumeInlineInstruction(line string) bool {
	// All instructions start from z
	if !strings.HasPrefix(line, "z") {
		return false
	}

	fields := strings.SplitN(line, " ", 2)
	key := fields[0]
	value := ""
	if len(fields) > 1 {
		value = strings.TrimSpace(fields[1])
	}

	switch key {
	case "zdel":
		n.Id = IdToRemove
	case "ztag":
		if value != "" {
			n.Tags = append(n.Tags, value)
		}
	case "zbook":
		n.Project = value
	default:
		return false
	}
	return true
}

func (n *Note) updateFromPropertiesLine(line string) {
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
	case "id":
		n.Id = value
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
		logrus.WithField("field", fields).Panic("unknown keys")
	}
}

func (n *Note) headerText() string {
	status := "active"
	if n.Archived {
		status = "archived"
	}
	return fmt.Sprintf(`<!--
znote: id=%s
znote: created=%s
znote: project=%s
znote: tags=%s
znote: status=%s
-->`, n.Id, n.CreationTime, n.Project, strings.Join(n.Tags, ","), status)
}

func (n *Note) contentToSave() string {
	return fmt.Sprintf("## %s\n%s\n%s\n", n.Title, n.headerText(), n.Content)
}
