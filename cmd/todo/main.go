package main

import (
	"flag"
	"fmt"
	"os"
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
	switch os.Args[1] {
	case "list":
		fmt.Println("Listing all tasks...")
		listTasks(data.Tasks)
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("t", "", "task title")
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse command failed!")
			return
		}
		if *title == "" {
			fmt.Println("Usage: todo add -t \"task title\"")
			return
		}
		data.Tasks = task.Add(*title, data.Tasks)
		err = store.WriteJSON(filePath, data)
		if err != nil {
			fmt.Println("Write file failed!")
			return
		}
		listTasks(data.Tasks)
	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		id := doneCmd.Int("id", 0, "task id")
		err := doneCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse command failed!")
			return
		}
		if *id == 0 {
			fmt.Println("Usage: todo done -id <task id>")
			return
		}
		data.Tasks, err = task.MarkDone(*id, data.Tasks)
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
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "task id")
		err := deleteCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse command failed!")
			return
		}
		if *id == 0 {
			fmt.Println("Usage: todo delete -id <task id>")
			return
		}
		data.Tasks, err = task.Delete(*id, data.Tasks)
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
	case "undo":
		undoCmd := flag.NewFlagSet("undo", flag.ExitOnError)
		id := undoCmd.Int("id", 0, "task id")
		err := undoCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse command failed!")
			return
		}
		if *id == 0 {
			fmt.Println("Usage: todo undo -id <task id>")
			return
		}
		data.Tasks, err = task.Undo(*id, data.Tasks)
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
		if taskData.Deleted {
			continue
		}
		status := ""
		if taskData.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, taskData.ID, taskData.Text)
	}
}
