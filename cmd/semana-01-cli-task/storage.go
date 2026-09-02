package main

import (
	"encoding/json"
	"os"
)

const fileName = "cmd/semana-01-cli-task/tasks.json"

// Recieves the list of tasks and saves it to disk in Json format
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
	//0644 is used for user read and write permissions
}

// Reads the json file from disk and converts it into a slice of Task
func LoadTasks() ([]Task, error) {
	// If the file dos not yet exist, return an empty slice without throwing and error
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
