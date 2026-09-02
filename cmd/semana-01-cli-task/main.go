package main

import (
	"fmt"
	"log"
)

func main() {
	tasks, err := LoadTasks()
	if err != nil {
		log.Fatalf("Error loading tasks: %v", err)
	}

	newTask := NewTask(len(tasks)+1, "Learn Golang")
	tasks = append(tasks, newTask)

	err = SaveTasks(tasks)

	if err != nil {
		log.Fatalf("Error saving tasks: %v", err)
	}

	fmt.Printf("Sucess! Total tasks saved %d\n", len(tasks))
}
