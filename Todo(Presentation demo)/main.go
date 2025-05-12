package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo/storage"
)

func printWelcome() {
	fmt.Println("Welcome to your TODO list.")
	fmt.Println("You can use the command add, delete, list or done.")
}

func main() {
	const fileName = "tasks.json"
	tasks, err := storage.LoadTasks(fileName)
	if err != nil {
		fmt.Println("Error", err)
		tasks = []storage.Task{}
	}

	printWelcome()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		args := strings.SplitN(input, " ", 2)
		cmd := args[0]

		switch cmd {
		case "add":
			if len(args) < 2 || strings.TrimSpace(args[1]) == "" {
				fmt.Println("Task name is required.")
				continue
			}
			tasks = storage.AddTask(tasks, args[1])
			storage.SaveTasks(fileName, tasks)
			fmt.Println("Task added:", args[1])

		case "list":
			storage.ListTasks(tasks)

		case "delete":
			if len(args) < 2 {
				fmt.Println("ID is required.")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}
			tasks = storage.DeleteTask(tasks, id)
			storage.SaveTasks(fileName, tasks)

		case "done":
			if len(args) < 2 {
				fmt.Println("ID is required.")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}
			tasks = storage.MarkDone(tasks, id)
			storage.SaveTasks(fileName, tasks)

		default:
			fmt.Println("Please enter a valid command.")
		}
	}
}
