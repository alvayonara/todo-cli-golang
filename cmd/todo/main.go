package main

import (
	"fmt"
	"os"
	"strconv"
	"todo-cli-golang/internal/store"
	"todo-cli-golang/internal/task"
	"todo-cli-golang/models"
)

const filePath = "datasource/tasks.json"

func main() {
	data, err := store.LoadJSON[models.Tasks](filePath)
	if err != nil {
		if os.IsNotExist(err) {
			data = models.Tasks{Tasks: []models.Task{}}
		} else {
			fmt.Println(err)
			return
		}
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		return
	}
	command := os.Args[1]
	switch command {
	case "list":
		fmt.Println("Listing all tasks...")
		listTasks(data.Tasks)
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Missing task name!")
			return
		}
		data.Tasks = task.Add(os.Args[2], data.Tasks)
		err := store.WriteJSON(filePath, data)
		if err != nil {
			fmt.Println("Write file failed!")
			return
		}
		listTasks(data.Tasks)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Missing task ID!")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		data.Tasks, err = task.MarkDone(id, data.Tasks)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = store.WriteJSON(filePath, data)
		if err != nil {
			fmt.Println("Write file failed!")
			return
		}
		listTasks(data.Tasks)
	default:
		fmt.Println("Unknown command")
	}
}

func listTasks(tasks []models.Task) {
	for _, taskData := range tasks {
		status := ""
		if taskData.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, taskData.ID, taskData.Text)
	}
}
