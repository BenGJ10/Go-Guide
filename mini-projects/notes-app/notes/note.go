package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const fileStorePath = "custom-notes/"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note *Note) Display() {
	fmt.Printf("\nTitle: %s\nContent: %s\nCreated At: %s\n", note.Title, note.Content, note.CreatedAt.Format(time.RFC1123))
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("Invalid! Title and content cannot be empty.")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		fmt.Println("Error converting note to JSON:", err)
		return err
	}

	return os.WriteFile(fileStorePath+fileName, json, 0644)
}
