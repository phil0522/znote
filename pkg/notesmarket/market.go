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
	Books     map[string]*Book
	PageBooks map[string]*PageBook
}

var (
	marketInstance *Market = nil
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

func (m *Market) GetOrCreatePageBook(bookName string) *PageBook {
	if bookName != "tech" {
		logrus.WithField("bookname", bookName).Panic("wrong page book name")
	}
	if book, ok := m.PageBooks[bookName]; ok {
		return book
	}
	logrus.Warnf("create new book %s", bookName)
	book := NewPageBook(bookName)
	m.PageBooks[bookName] = book
	return book
}

func (m *Market) SaveAll() {
	for _, book := range m.Books {
		book.saveToDisk()
	}
}

func (m *Market) loadAll() {
	logrus.WithField("root", RootDir).Warn("load whole note market")
	err := filepath.Walk(RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Fatalf("error reading file %s", path)
			return nil
		}

		if info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		bookName := getBookName(path)
		if bookName == "" || bookName == "tmp" {
			return nil
		} else if bookName == "tech" {
			_ = m.GetOrCreatePageBook(bookName)
		} else if isDataFile(path, info) {
			book := m.GetOrCreateBook(bookName)
			nf := NewNoteFile(path)
			nf.load()
			book.noteFiles[path] = nf
			book.Notes.mergeNoteSet(nf.notes)
		}

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
	if marketInstance == nil {
		marketInstance = &Market{
			Books:     make(map[string]*Book),
			PageBooks: make(map[string]*PageBook),
		}
		marketInstance.loadAll()
	}
	return marketInstance
}

func (m *Market) Reload() {
	for k := range m.Books {
		delete(m.Books, k)
	}
	for k := range m.PageBooks {
		delete(m.PageBooks, k)
	}
	m.loadAll()
}
