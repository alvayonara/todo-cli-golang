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
	tests := []struct {
		name      string
		id        int
		wantDone  bool
		wantError bool
	}{
		{"mark existing task", 1, true, false},
		{"task already done", 2, true, true},
		{"task not found", 99, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks := []models.Task{
				{ID: 1, Text: "A", Done: false},
				{ID: 2, Text: "B", Done: true},
			}

			updated, err := MarkDone(tt.id, tasks)
			if tt.wantError && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tt.wantError && updated[tt.id-1].Done != tt.wantDone {
				t.Errorf("expected Done=%v", tt.wantDone)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tasks := []models.Task{{ID: 1, Text: "A"}}
	updated, err := Delete(1, tasks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !updated[0].Deleted {
		t.Errorf("task should be deleted")
	}
}

func TestUndo(t *testing.T) {
	tasks := []models.Task{{ID: 1, Text: "A", Deleted: true}}
	updated, err := Undo(1, tasks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated[0].Deleted {
		t.Errorf("task should be restored")
	}
}
