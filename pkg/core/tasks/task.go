package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Task struct {
	Name        string
	Description string
}

func input(text string) string {
	// This will read the whole line
	// including whitespaces, until is inserted a new line
	reader := bufio.NewReader(os.Stdin)

	var inputStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#76946A")).
		Foreground(lipgloss.Color("#76946A"))

	var promptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#76946A"))

	fmt.Println(inputStyle.Render("Enter the task " + text))
	fmt.Print(promptStyle.Render("> "))
	property, err := reader.ReadString('\n')
	property = strings.TrimSpace(property)

	if err != nil {
		fmt.Printf("Error reading task %v: ", text)
		fmt.Println(err)
		os.Exit(1)
	}

	return property
}

func newTask() Task {
	taskName := input("name")
	taskDescription := input("description")

	return Task{
		Name:        taskName,
		Description: taskDescription,
	}
}

func PrintTask() {
	task := newTask()

	fmt.Println()

	var (
		headerStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#76946A")).
				Bold(true).
				Align(lipgloss.Center)

		cellStyle = lipgloss.NewStyle().
				Padding(0, 1).
				Width(14)
	)

	rows := [][]string{
		{task.Name, task.Description},
	}

	table := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#76946A"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case table.HeaderRow:
				return headerStyle
			default:
				return cellStyle
			}
		}).
		Headers("NAME", "DESCRIPTION").
		Rows(rows...)

	var congratScreen = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#76946A")).
		Foreground(lipgloss.Color("#76946A"))

	fmt.Println(congratScreen.Render("TASK CREATED!"))

	fmt.Println(table)
}
