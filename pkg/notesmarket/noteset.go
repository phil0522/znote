package notesmarket

import (
	"sort"

	"github.com/sirupsen/logrus"
)

type NoteSet struct {
	notes map[string]*Note
}

func NewNoteSet() NoteSet {
	return NoteSet{
		notes: make(map[string]*Note),
	}
}

func (ns *NoteSet) HasNoteById(key string) bool {
	_, ok := ns.notes[key]
	return ok
}

func (ns *NoteSet) mergeNoteSet(other NoteSet) {
	for k, v := range other.notes {
		ns.notes[k] = v
	}
}

func (ns *NoteSet) mergeNotes(notes []*Note) {
	for _, n := range notes {
		ns.notes[n.Id] = n
	}
}

func (ns *NoteSet) removeNoteByKey(key string) {
	delete(ns.notes, key)
}

func (ns *NoteSet) upsertNode(note Note) {
	if note.Id == "" {
		logrus.WithField("note", note).Panic("Empty Note Key")
	}
	ns.notes[note.Id] = &note
}

func (ns *NoteSet) ToOrderedList() []*Note {
	r := make([]*Note, 0, len(ns.notes))
	keys := make([]string, 0, len(ns.notes))

	for k := range ns.notes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	i := len(keys) - 1
	for i >= 0 {
		r = append(r, ns.notes[keys[i]])
		i -= 1
	}
	return r
}

func (ns *NoteSet) GetNote(key string) Note {
	v, ok := ns.notes[key]
	if ok {
		return *v
	}
	logrus.WithField("note", key).Fatal("can not find note")
	return EmptyNote()
}
