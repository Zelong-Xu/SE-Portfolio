package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo/storage"
)

// Go uses package-based modularity (Package-based Modularity),
// making it easy to reuse and organize code across files and folders.

func printUsage() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  add <task name>      → Add a new task")
	fmt.Println("  list                 → List all tasks")
	fmt.Println("  done <task_id>       → Mark task as done/undone")
	fmt.Println("  delete <task_id>     → Delete a task")
	fmt.Println("  exit                 → Exit the program")
}

func main() {
	const fileName = "tasks.json"

	tasks, err := storage.LoadTasks(fileName)
	if err != nil {
		fmt.Println("❌ Failed to load tasks:", err)
		tasks = []storage.Task{}
	}
	// Go’s error handling is explicit and straightforward,
	// using multiple return values to indicate success or failure.

	fmt.Println("📝 Welcome to the Go Todo List App!")
	printUsage()

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
		// Go’s switch is more readable than long if-else chains,
		// and does not fall through by default—safer and cleaner.

		case "add":
			if len(args) < 2 || strings.TrimSpace(args[1]) == "" {
				fmt.Println("⚠️  Task name is required.")
				continue
			}
			tasks = storage.AddTask(tasks, args[1])
			storage.SaveTasks(fileName, tasks)
			fmt.Println("✅ Task added:", args[1])

		case "list":
			storage.ListTasks(tasks)

		case "done":
			if len(args) < 2 {
				fmt.Println("⚠️  Task ID is required.")
				continue
			}
			id, err := strconv.Atoi(args[1])
			// Go is a statically typed language (Statically Typed Language),
			// ensuring type safety and early error detection at compile time.

			if err != nil {
				fmt.Println("⚠️  Invalid ID.")
				continue
			}
			tasks = storage.MarkDone(tasks, id)
			storage.SaveTasks(fileName, tasks)

		case "delete":
			if len(args) < 2 {
				fmt.Println("⚠️  Task ID is required.")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("⚠️  Invalid ID.")
				continue
			}
			tasks = storage.DeleteTask(tasks, id)
			storage.SaveTasks(fileName, tasks)

		default:
			fmt.Println("⚠️  Unknown command.")
			printUsage()
		}
	}
}
