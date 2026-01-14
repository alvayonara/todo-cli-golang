package main

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-cli-golang/models"
)

func main() {
	var filePath = "datasource/tasks.json"
	data, err := LoadJSON[models.Tasks](filePath)
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
		data.Tasks = addTask(os.Args[2], data.Tasks)
		err := writeJSON(filePath, data)
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
	for _, task := range tasks {
		status := ""
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Text)
	}
}

func addTask(taskTitle string, tasks []models.Task) []models.Task {
	task := models.Task{ID: generateId(tasks), Text: taskTitle, Done: false}
	tasks = append(tasks, task)
	fmt.Printf("Added new task: %s succeed!\n", taskTitle)
	return tasks
}

func LoadJSON[T any](filename string) (T, error) {
	var data T
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(fileData, &data)
}

func writeJSON(filename string, tasks models.Tasks) error {
	updatedData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}
	err = os.WriteFile(filename, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}
	return nil
}

func generateId(tasks []models.Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}
