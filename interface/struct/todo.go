package todoStruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("text is required")
	}
	return &Todo{Text: text}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) SaveToFile() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	return os.WriteFile(fileName, json, 0644)
}
