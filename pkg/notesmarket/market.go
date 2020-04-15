package notesmarket

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// Market is the manager of the whole notes system. A market may contains
// multiple books, and each book may contains multiple notes.
// Market is also responsible to serialize all notes to files and load them.
type Market struct {
	noteFiles map[string]*NoteFile
	Books     map[string]*Book
}

var (
	marketInstance = &Market{
		noteFiles: make(map[string]*NoteFile),
		Books:     make(map[string]*Book),
	}
)

func (m *Market) GetOrCreateBook(bookName string) *Book {
	if book, ok := m.Books[bookName]; ok {
		return book
	}
	book := NewBook()
	book.Name = bookName
	m.Books[book.Name] = book
	return book
}
func (m *Market) SaveAll() {
	for _, book := range m.Books {
		book.saveToDisk()
	}
}

func init() {
	marketInstance.loadAll()
}

func (m *Market) loadAll() {
	logrus.WithField("root", RootDir).Debug("load whole note market")
	err := filepath.Walk(RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Fatalf("error reading file %s", path)
			return nil
		}
		if !isDataFile(path, info) {
			return nil
		}

		nf := NewNoteFile(path)
		nf.load()

		bookName := getBookName(path)
		if bookName == "" {
			return nil
		}
		book := m.GetOrCreateBook(bookName)

		book.noteFiles[path] = nf
		book.Notes.mergeNoteSet(nf.notes)
		return nil
	})

	if err != nil {
		logrus.Fatalf("failed to walk")
	}
}

var (
	validBaseName = regexp.MustCompile(`^[0-9]*.md$`)
)

func isDataFile(path string, info os.FileInfo) bool {
	if info.IsDir() {
		logrus.Debugf("directory %s is a directory, skipping", path)
		return false
	}

	if validBaseName.MatchString(info.Name()) {
		return true
	}

	logrus.WithField("path", path).Debug("Not an expected data file")
	return false
}

func getBookName(path string) string {
	relativePath, err := filepath.Rel(RootDir, path)
	if err != nil {
		logrus.WithField("path", path).Fatal("Failed to calculate relative path")
		return ""
	}

	fileSegments := strings.Split(relativePath, string(filepath.Separator))
	if len(fileSegments) == 1 {
		logrus.WithField("path", path).Fatal("file is under the root, no book assigned")
		return ""
	}
	return fileSegments[0]
}

func GetNotesMarket() *Market {
	return marketInstance
}
