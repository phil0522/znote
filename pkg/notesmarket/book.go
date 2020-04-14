package notesmarket

import "github.com/sirupsen/logrus"

// Book may contains multiple notes. Each book has its own directory.
// Notes may be duplicated to different locations.
// ByProject, ByTags, ByMonth
type Book struct {
	Name  string
	Notes NoteSet

	noteFiles map[string]*NoteFile
}

func NewBook() *Book {
	return &Book{
		Notes:     NewNoteSet(),
		noteFiles: make(map[string]*NoteFile),
	}
}
func (b *Book) RemoveNote(n Note) {

}
func (b *Book) AddNote(n Note) {
	logrus.WithField("book", b.Name).WithField("note", n).Info("Add new Note")
	notePath := b.getFilePathFromCreationTime(n.CreationTime)
	nf, ok := b.noteFiles[notePath]
	if !ok {
		nf = NewNoteFile(notePath)
	}
	nf.notes.upsertNode(n)
	nf.changed = true

	b.noteFiles[notePath] = nf
}

func (b *Book) saveToDisk() {
	for _, nf := range b.noteFiles {
		nf.save()
	}
}
