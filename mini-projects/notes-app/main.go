package main

import (
	"bufio"
	"fmt"
	"notes-app/notes"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := notes.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully!")

}

func getNoteData() (string, string) {
	title := getUserInput("Enter the note title: ")
	content := getUserInput("Enter the note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	return input
}
