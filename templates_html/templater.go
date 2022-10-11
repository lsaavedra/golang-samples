package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/template"
)

type BoardTasks struct {
	TasksCount int
	ToDoTasks  []string
}

func errorCheck(err error) {
	// Handle errors
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Get our text from the file
	taskVals := getStrings("tasks.txt")
	fmt.Println("tasks from file ", taskVals)

	// Print to the terminal
	fmt.Printf("%#v\n", taskVals)
	// Create a template using the html
	tmpl, err := template.ParseFiles("board_tasks.html")
	errorCheck(err)

	// Create a todo list with the number
	tasks := BoardTasks{
		TasksCount: len(taskVals),
		ToDoTasks:  taskVals,
	}

	// Write the template to the ResponseWriter
	// Pass the todo struct data
	writer := os.Stdout
	err = tmpl.Execute(writer, tasks)
}

// Retreives lines of text from a file
func getStrings(fileName string) []string {
	var lines []string

	// Try to open the file (It must exist)
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	// Close file when the function ends
	defer file.Close()

	// Read lines of text and save to lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())

	// Return the text
	return lines
}
