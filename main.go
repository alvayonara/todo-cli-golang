package main

import (
	"fmt"
	"os"
)

func main() {
	taskList := []Task{
		{ID: 1, Text: "Hello World", Done: true},
		{ID: 2, Text: "Do something", Done: false},
		{ID: 3, Text: "Cooking", Done: true},
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		return
	}
	command := os.Args[1]
	switch command {
	case "list":
		fmt.Println("Listing all tasks...")
		listTask(taskList)
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Missing task name!")
			return
		}
		taskList = addTask(os.Args[2], taskList)
		listTask(taskList)
	default:
		fmt.Println("Unknown command")
	}
}

func listTask(tasks []Task) {
	for _, task := range tasks {
		status := ""
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Text)
	}
}

func addTask(taskTitle string, tasks []Task) []Task {
	task := Task{ID: len(tasks) + 1, Text: taskTitle, Done: false}
	tasks = append(tasks, task)
	fmt.Printf("Added new task: %s succeed!\n", taskTitle)
	return tasks
}

type Task struct {
	ID   int
	Text string
	Done bool
}
