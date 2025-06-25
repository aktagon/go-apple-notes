package notes

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

//go:embed jxa/create_note.jxa
var createNoteScript string

//go:embed jxa/get_all_notes.jxa
var getAllNotesScript string

//go:embed jxa/find_notes.jxa
var findNotesScript string

//go:embed jxa/update_note.jxa
var updateNoteScript string

//go:embed jxa/delete_note.jxa
var deleteNoteScript string

//go:embed jxa/get_note_by_id.jxa
var getNoteByIDScript string

type Client struct{}

// escapeJXAString escapes special characters in strings to prevent JXA injection
func escapeJXAString(s string) string {
	replacer := strings.NewReplacer(
		`\`, `\\`,
		`"`, `\"`,
		"\n", `\n`,
		"\r", `\r`,
		"\t", `\t`,
	)
	return replacer.Replace(s)
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) CreateNote(title, body, folderName string) (*CreateNoteResult, error) {
	if folderName == "" {
		folderName = "Notes"
	}

	tmpl, err := template.New("create_note").Parse(createNoteScript)
	if err != nil {
		return &CreateNoteResult{
			Success: false,
			Message: fmt.Sprintf("Failed to parse template: %v", err),
		}, nil
	}

	var scriptBuilder strings.Builder
	err = tmpl.Execute(&scriptBuilder, struct {
		Title      string
		Body       string
		FolderName string
	}{
		Title:      title,
		Body:       body,
		FolderName: folderName,
	})
	if err != nil {
		return &CreateNoteResult{
			Success: false,
			Message: fmt.Sprintf("Failed to execute template: %v", err),
		}, nil
	}

	script := scriptBuilder.String()

	output, err := executeJXA(script)
	if err != nil {
		return &CreateNoteResult{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	var result CreateNoteResult
	if err := json.Unmarshal(output, &result); err != nil {
		return &CreateNoteResult{
			Success: false,
			Message: fmt.Sprintf("Failed to parse result: %v", err),
		}, nil
	}

	return &result, nil
}

func (c *Client) GetAllNotes() ([]Note, error) {
	output, err := executeJXA(getAllNotesScript)
	if err != nil {
		return nil, err
	}

	var notes []Note
	if err := json.Unmarshal(output, &notes); err != nil {
		return nil, fmt.Errorf("failed to parse notes JSON: %w", err)
	}

	return notes, nil
}

func (c *Client) FindNotes(searchText string) ([]Note, error) {
	tmpl, err := template.New("find_notes").Parse(findNotesScript)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	var scriptBuilder strings.Builder
	err = tmpl.Execute(&scriptBuilder, struct {
		SearchText string
	}{
		SearchText: searchText,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	output, err := executeJXA(scriptBuilder.String())
	if err != nil {
		return nil, err
	}

	var notes []Note
	if err := json.Unmarshal(output, &notes); err != nil {
		return nil, fmt.Errorf("failed to parse search results JSON: %w", err)
	}

	if len(notes) == 0 {
		allNotes, err := c.GetAllNotes()
		if err != nil {
			return nil, err
		}

		searchLower := strings.ToLower(searchText)
		for _, note := range allNotes {
			if strings.Contains(strings.ToLower(note.Name), searchLower) ||
				strings.Contains(strings.ToLower(note.Content), searchLower) {
				notes = append(notes, note)
			}
		}
	}

	return notes, nil
}

func (c *Client) UpdateNote(noteID, newTitle, newContent string) error {
	tmpl, err := template.New("update_note").Parse(updateNoteScript)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var scriptBuilder strings.Builder
	err = tmpl.Execute(&scriptBuilder, struct {
		NoteID     string
		NewTitle   string
		NewContent string
	}{
		NoteID:     noteID,
		NewTitle:   newTitle,
		NewContent: newContent,
	})
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	_, err = executeJXA(scriptBuilder.String())
	return err
}

func (c *Client) DeleteNote(noteID string) error {
	tmpl, err := template.New("delete_note").Parse(deleteNoteScript)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var scriptBuilder strings.Builder
	err = tmpl.Execute(&scriptBuilder, struct {
		NoteID string
	}{
		NoteID: noteID,
	})
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	_, err = executeJXA(scriptBuilder.String())
	return err
}

func (c *Client) GetNoteByID(noteID string) (*Note, error) {
	tmpl, err := template.New("get_note_by_id").Parse(getNoteByIDScript)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	var scriptBuilder strings.Builder
	err = tmpl.Execute(&scriptBuilder, struct {
		NoteID string
	}{
		NoteID: noteID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	output, err := executeJXA(scriptBuilder.String())
	if err != nil {
		return nil, err
	}

	var note Note
	if err := json.Unmarshal(output, &note); err != nil {
		return nil, fmt.Errorf("failed to parse note JSON: %w", err)
	}

	return &note, nil
}
