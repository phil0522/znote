package notesmarket

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// Market is the manager of the whole notes system. A market may contains
// multiple books, and each book may contains multiple notes.
// Market is also responsible to serialize all notes to files and load them.
type Market struct {
	noteFiles map[string]noteFile
	Books     map[string]*Book
}

var (
	marketInstance = &Market{
		noteFiles: make(map[string]noteFile),
		Books:     make(map[string]*Book),
	}
)

func (m *Market) getOrCreateBook(bookName string) *Book {
	if book, ok := m.Books[bookName]; ok {
		return book
	}
	book := &Book{}
	book.Name = bookName
	m.Books[book.Name] = book
	return book
}
func (m *Market) SaveAll() {

}

func (m *Market) Load() {

}

func init() {
	marketInstance.loadAll()
}

func (m *Market) loadAll() {
	logrus.WithField("root", RootDir).Warn("load whole note market")
	err := filepath.Walk(RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Fatalf("error reading file %s", path)
			return nil
		}
		if info.IsDir() {
			logrus.Debugf("directory %s is a directory, skipping", path)
			return nil
		}

		if strings.HasSuffix(path, "README.md") || !strings.HasSuffix(path, ".md") {
			logrus.WithField("path", path).Debug("Skip non markdown file")
			return nil
		}

		relativePath, err := filepath.Rel(RootDir, path)
		if err != nil {
			logrus.WithField("base", RootDir).WithField("path", path).Fatal("calculate relative path failed")
			return nil
		}
		fileSegments := strings.Split(relativePath, string(filepath.Separator))
		if len(fileSegments) == 1 {
			logrus.WithField("path", fileSegments).Info("skip file under root")
		}

		nf := noteFile{}
		nf.load(path)
		m.noteFiles[path] = nf

		book := m.getOrCreateBook(fileSegments[0])
		book.Notes = append(book.Notes, nf.notes...)
		return nil
	})

	if err != nil {
		logrus.Fatalf("failed to walk")
	}
}

func GetNotesMarket() *Market {
	return marketInstance
}
