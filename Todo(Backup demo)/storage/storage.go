package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}, nil
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(filename string, tasks []Task) {

	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(filename, data, 0644)
}

func AddTask(tasks []Task, title string) []Task {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:    id,
		Title: title,
		Done:  false,
	}
	return append(tasks, newTask)
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks found.")
		return
	}

	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}
}

func MarkDone(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Printf("✅ Task %d marked as done.\n", id)
		}
	}
	return tasks
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			fmt.Printf("🗑️ Task %d deleted.\n", id)

			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	fmt.Println("❌ Task ID not found.")
	return tasks
}
