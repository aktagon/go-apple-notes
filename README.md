# Go Apple Notes

A Go library for CRUD operations on Apple Notes using JavaScript for Automation (JXA).

## Requirements

- macOS 10.10+ (Yosemite or later)
- Go 1.19+
- Apple Notes app

## Installation

```bash
go get github.com/aktagon/go-apple-notes
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/aktagon/go-apple-notes/pkg/notes"
)

func main() {
    client := notes.NewClient()

    // Create a note
    result, err := client.CreateNote(
        "My Note",
        "# Hello World\n\nThis is a test note.",
        "Notes", // folder name (optional)
    )
    if err != nil {
        log.Fatal(err)
    }

    if result.Success {
        fmt.Printf("Created note: %s (ID: %s)\n",
            result.Note.Name, result.Note.ID)
    }

    // List all notes
    allNotes, err := client.GetAllNotes()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Total notes: %d\n", len(allNotes))

    // Search notes
    foundNotes, err := client.FindNotes("Hello World")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d matching notes\n", len(foundNotes))
}
```

### API Methods

- `CreateNote(title, content, folder string) (*CreateNoteResult, error)`
- `GetAllNotes() ([]Note, error)`
- `FindNotes(searchText string) ([]Note, error)`
- `GetNoteByID(noteID string) (*Note, error)`
- `UpdateNote(noteID, title, content string) error`
- `DeleteNote(noteID string) error`

## CLI Tool

Build and use the included CLI tool:

```bash
make build
./bin/notes-cli create "My Note" "Content here" "Folder"
./bin/notes-cli list
./bin/notes-cli search "keyword"
```

## Development

```bash
# Run tests
make test

# Run example
make run-example

# Format code
make fmt

# Clean build artifacts
make clean
```

## Permissions

On first use, macOS will prompt for permission to access the Notes app. Grant access to enable the library to function.

