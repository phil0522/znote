package notesmarket

// Note is the basic unit of the notes system.
type Note struct {
	Title        string
	Content      string
	Project      string
	Tags         []string
	Properties   map[string]string
	CreationDate string
	Archived     bool
}
