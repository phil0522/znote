package notesmarket

// Book may contains multiple notes. Each book has its own directory.
// Notes may be duplicated to different locations.
// ByProject, ByTags, ByMonth
type Book struct {
	Name  string
	Notes []Note

	ByTags    map[string]*Note
	ByProject map[string]*Note
	ByMonth   map[string]*Note
}
