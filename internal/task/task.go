package task

import (
	"fmt"
	"todo-cli-golang/models"
)

func Add(taskTitle string, tasks []models.Task) []models.Task {
	task := models.Task{ID: generateId(tasks), Text: taskTitle, Done: false}
	tasks = append(tasks, task)
	fmt.Printf("Added new task: %s succeed!\n", taskTitle)
	return tasks
}

func MarkDone(id int, tasks []models.Task) ([]models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			if task.Done {
				return tasks, fmt.Errorf("task %d is already done", id)
			}
			tasks[i].Done = true
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("task %d not found", id)
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
