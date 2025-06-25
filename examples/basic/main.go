package main

import (
	"fmt"
	"log"

	"github.com/aktagon/go-apple-notes/pkg/notes"
)

func main() {
	client := notes.NewClient()

	fmt.Println("Creating a new note...")
	result, err := client.CreateNote(
		"My Test Note",
		"# Hello World\n\nThis is a test note created from Go.\n\n- Item 1\n- Item 2",
		"Notes",
	)
	if err != nil {
		log.Fatal(err)
	}

	if result.Success {
		fmt.Printf("Created note: %s\n", result.Note.Name)
		fmt.Printf("  ID: %s\n", result.Note.ID)
		fmt.Printf("  Folder: %s\n", result.FolderName)
		if result.UsedDefaultFolder {
			fmt.Println("  (Created default Notes folder)")
		}

		fmt.Println("\nSearching for notes...")
		foundNotes, err := client.FindNotes("Test Note")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found %d notes matching 'Test Note'\n", len(foundNotes))

		fmt.Println("\nListing all notes...")
		allNotes, err := client.GetAllNotes()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Total notes: %d\n", len(allNotes))
		for i, note := range allNotes {
			if i < 5 {
				fmt.Printf("  - %s (Folder: %s)\n", note.Name, note.Folder)
			}
		}
		if len(allNotes) > 5 {
			fmt.Printf("  ... and %d more\n", len(allNotes)-5)
		}
	} else {
		fmt.Printf("Failed to create note: %s\n", result.Message)
	}
}
