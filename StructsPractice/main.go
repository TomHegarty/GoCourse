package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	//fetch user input
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	printSomething("cat")
	printSomething(1)
	printSomething(1.0)
	printSomething(userNote.Save())

	// number := add(3, 4)
	// number2 := add(3.34, 4.43)
	// string := add("3.34", "4.43")
}

func printSomething(value interface{}) {

	typedValue, ok := value.(int)

	if !ok {
		fmt.Println(" typedValue not an int ")
		return
	} else {
		fmt.Println(" int is: ", typedValue)
	}

	switch value.(type) {
	case string:
		fmt.Println("string: ", value)
	case int:
		fmt.Println("int: ", value)
	case float64:
		fmt.Println("float64: ", value)
	default:
		fmt.Println("unprintable type ")
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // to handle windows

	return text
}

func add[T int | float64 | float32 | string](a, b T) T {
	return a + b
}
