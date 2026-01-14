package store

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-cli-golang/models"
)

func LoadJSON[T any](filename string) (T, error) {
	var data T
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(fileData, &data)
}

func WriteJSON(filename string, tasks models.Tasks) error {
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
