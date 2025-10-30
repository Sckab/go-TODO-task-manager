package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Name        string
	Description string
}

func newTask() Task {
	// This will read the whole line
	// including whitespaces, until is inserted a new line
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the task name: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	fmt.Print("Enter the task description: ")
	taskDescription, _ := reader.ReadString('\n')
	taskDescription = strings.TrimSpace(taskDescription)

	return Task{
		Name:        taskName,
		Description: taskDescription,
	}
}

func PrintTask() {
	task := newTask()
	fmt.Println()
	fmt.Println("Task created!")
	fmt.Println()
	fmt.Printf("Task name: %v \n", task.Name)
	fmt.Printf("Task description: %v \n", task.Description)
}
