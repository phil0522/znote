package notesmarket

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

type PageBook struct {
	Name string
	// abs path to title mapping
	Pages map[string]string
}

func NewPageBook(name string) *PageBook {
	return &PageBook{
		Name:  name,
		Pages: make(map[string]string),
	}
}

func (pb *PageBook) ReloadAll() {
	pb.Pages = make(map[string]string)
	logrus.WithField("root", RootDir).Debug("load page book")
	bookRoot := filepath.Join(RootDir, pb.Name)
	err := filepath.Walk(bookRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.WithError(err).WithField("path", path).Fatal("error reading file")
			return nil
		}

		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		pb.loadPage(path)
		return nil
	})

	if err != nil {
		logrus.WithError(err).Panic("failed to walk")
	}
}

func (pb *PageBook) loadPage(path string) {
	file, err := os.Open(path)
	if err != nil {
		logrus.WithError(err).WithField("path", path).Warn("failed to open file")
		pb.Pages[path] = "Failed to open file"
		return
	}
	buf := make([]byte, 300)
	n, err := file.Read(buf)
	if err != nil {
		logrus.WithError(err).WithField("path", path).Warn("failed to read file")
		pb.Pages[path] = "Failed to read file"
		return
	}
	lines := strings.Split(string(buf[0:n]), "\n")
	title := strings.TrimLeft(lines[0], "#")

	pb.Pages[path] = title

	logrus.Warnf("load new book %s: %s", path, title)
}

func (pb *PageBook) UpdateToc() {
	logrus.WithField("root", RootDir).Debug("Update toc")

	keys := make([]string, 0, len(pb.Pages))
	for key := range pb.Pages {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.TrimSuffix(keys[i], "README.md") < strings.TrimSuffix(keys[j], "README.md")
	})

	lastDir := ""
	f, err := os.Create(filepath.Join(RootDir, pb.Name, pb.Name+".toc"))

	if err != nil {
		logrus.WithError(err).Panic("failed to update book toc")
	}

	defer f.Close()

	for _, path := range keys {
		relPath, _ := filepath.Rel(filepath.Join(RootDir, pb.Name), path)
		baseDir, _ := filepath.Split(relPath)
		if baseDir != lastDir {
			if !strings.HasSuffix(relPath, "README.md") {
				fmt.Fprintf(f, "%s,\n", baseDir)
			}
			fmt.Fprintf(f, "%s,%s\n", relPath, pb.Pages[path])
			lastDir = baseDir
		} else {
			fmt.Fprintf(f, "%s,%s\n", relPath, pb.Pages[path])
		}
	}
}
