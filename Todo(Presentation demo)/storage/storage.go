package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func AddTask(tasks []Task, title string) []Task {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	newTask := Task{
		ID:   id,
		Name: title,
		Done: false,
	}
	return append(tasks, newTask)
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[X]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Name)
	}
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			fmt.Printf("Task %d deleted.\n", id)
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	fmt.Println("Task ID not found.")
	return tasks
}

func SaveTasks(filename string, tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}, nil
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func MarkDone(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Printf("Task %d is done.\n", id)
			break
		}
	}
	return tasks
}
