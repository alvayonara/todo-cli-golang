package task

import (
	"testing"
	"todo-cli-golang/models"
)

func TestAdd(t *testing.T) {
	tasks := []models.Task{}
	tasks = Add("Playing FIFA", tasks)
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Text != "Playing FIFA" {
		t.Errorf("expected task text Playing FIFA, got %q", tasks[0].Text)
	}
	if tasks[0].Done {
		t.Errorf("new task should not be done")
	}
}

func TestMarkDone(t *testing.T) {
	tasks := []models.Task{
		{1, "Playing Fifa", false, false},
	}
	tasks, _ = MarkDone(1, tasks)
	if tasks[0].Done != true {
		t.Errorf("expected task done Playing FIFA, got %t", tasks[0].Done)
	}
}
