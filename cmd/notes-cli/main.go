package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aktagon/go-apple-notes/pkg/notes"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: notes-cli <command> [args...]")
		fmt.Println("Commands:")
		fmt.Println("  create <title> <content> [folder]")
		fmt.Println("  list")
		fmt.Println("  search <text>")
		fmt.Println("  get <id>")
		fmt.Println("  update <id> <title> <content>")
		fmt.Println("  delete <id>")
		os.Exit(1)
	}

	client := notes.NewClient()
	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 4 {
			log.Fatal("Usage: notes-cli create <title> <content> [folder]")
		}
		title := os.Args[2]
		content := os.Args[3]
		folder := ""
		if len(os.Args) > 4 {
			folder = os.Args[4]
		}

		result, err := client.CreateNote(title, content, folder)
		if err != nil {
			log.Fatal(err)
		}
		if result.Success {
			fmt.Printf("Created note: %s (ID: %s) in folder: %s\n",
				result.Note.Name, result.Note.ID, result.FolderName)
		} else {
			fmt.Printf("Failed to create note: %s\n", result.Message)
		}

	case "list":
		allNotes, err := client.GetAllNotes()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found %d notes:\n", len(allNotes))
		for _, note := range allNotes {
			fmt.Printf("- %s (ID: %s, Folder: %s)\n", note.Name, note.ID, note.Folder)
		}

	case "search":
		if len(os.Args) < 3 {
			log.Fatal("Usage: notes-cli search <text>")
		}
		searchText := os.Args[2]
		foundNotes, err := client.FindNotes(searchText)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found %d notes matching '%s':\n", len(foundNotes), searchText)
		for _, note := range foundNotes {
			fmt.Printf("- %s (ID: %s, Folder: %s)\n", note.Name, note.ID, note.Folder)
		}

	case "get":
		if len(os.Args) < 3 {
			log.Fatal("Usage: notes-cli get <id>")
		}
		noteID := os.Args[2]
		note, err := client.GetNoteByID(noteID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Title: %s\n", note.Name)
		fmt.Printf("Folder: %s\n", note.Folder)
		fmt.Printf("Content:\n%s\n", note.Content)

	case "update":
		if len(os.Args) < 5 {
			log.Fatal("Usage: notes-cli update <id> <title> <content>")
		}
		noteID := os.Args[2]
		newTitle := os.Args[3]
		newContent := os.Args[4]
		err := client.UpdateNote(noteID, newTitle, newContent)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated note with ID: %s\n", noteID)

	case "delete":
		if len(os.Args) < 3 {
			log.Fatal("Usage: notes-cli delete <id>")
		}
		noteID := os.Args[2]
		err := client.DeleteNote(noteID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted note with ID: %s\n", noteID)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
