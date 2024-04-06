package main

import (
	"bufio"
	"errors"
	"fmt"
	todoStruct "github.com/SJ22032003/go-interface/struct"
	"os"
	"strings"
)

type genericUtil interface {
	SaveToFile() error
	Display()
}

func main() {

	text, err := readTodo("Enter a todo:")
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todoStruct.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)

}

func outputData(data genericUtil) {
	data.Display()
	saveErr := saveToFile(data)
	if saveErr != nil {
		fmt.Println(saveErr)
		return
	}
}

func saveToFile(data genericUtil) error {
	return data.SaveToFile()
}

func readTodo(prompt string) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", errors.New("failed to read text")
	}

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "") // for windows

	return text, nil
}
