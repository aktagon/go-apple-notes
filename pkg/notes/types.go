package notes

type Note struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Folder  string `json:"folder"`
}

type CreateNoteResult struct {
	Success           bool   `json:"success"`
	Note              *Note  `json:"note,omitempty"`
	Message           string `json:"message,omitempty"`
	FolderName        string `json:"folderName,omitempty"`
	UsedDefaultFolder bool   `json:"usedDefaultFolder"`
}

type NotesClient interface {
	CreateNote(title, content, folder string) (*CreateNoteResult, error)
	GetAllNotes() ([]Note, error)
	FindNotes(searchText string) ([]Note, error)
	GetNoteByID(noteID string) (*Note, error)
	UpdateNote(noteID, title, content string) error
	DeleteNote(noteID string) error
}
