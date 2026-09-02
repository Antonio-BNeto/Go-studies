package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: go run ./cmd/semana-01-cli-task [add|list|complete] [Args]")
	}

	tasks, err := LoadTasks()

	if err != nil {
		log.Fatalf("Error loading tasks: %v", err)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please, Add the task title")
			return
		}
		title := os.Args[2]
		id := len(tasks) + 1
		newTask := NewTask(id, title)
		tasks = append(tasks, newTask)
		err = SaveTasks(tasks)
		if err != nil {
			log.Fatalf("Error saving tasks: %v", err)
		}
		fmt.Printf("Successfully added task: %v, ID: %d\n", title, id)
	case "list":
		if len(tasks) == 0 {
			fmt.Println("The task list is empty.")
			return
		}
		fmt.Println("=========Your Tasks=========")

		for _, t := range tasks {
			status := "[ ]"
			if t.Completed {
				status = "[X]"
			}
			fmt.Printf("%d, %s %s\n", t.ID, status, t.Title)
		}

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please, Add the task title")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Please, use the Integer number")
			return
		}

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Completed = true
				found = true
				break
			}
		}

		if found != true {
			fmt.Println("The tasks not found")
			return
		}

		err = SaveTasks(tasks)
		if err != nil {
			log.Fatalf("Error updating task", err)
		}

		fmt.Println("Task successfully updated")
	default:
		fmt.Println("Command not exist")
	}
}
