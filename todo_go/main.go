package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	loadTaskData()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nEnter a number to choose:")
		fmt.Println("1. Add a task")
		fmt.Println("2. Remove a task")
		fmt.Println("3. List tasks")
		fmt.Println("4. Toggle task status (done/undone)")
		fmt.Println("0. Exit")
		fmt.Print("> ")
		choiceRaw, _ := reader.ReadString('\n')
		choiceRaw = strings.TrimSpace(choiceRaw)
		choice, err := strconv.Atoi(choiceRaw)
		if err != nil {
			fmt.Println("Invalid input. Try again!")
			continue
		}
		switch choice {
		case 1:
			for {
				fmt.Print("Add a task: ")
				taskRaw, _ := reader.ReadString('\n')
				taskRaw = strings.TrimSpace(taskRaw)
				taskNorm := normalizeSpaces(taskRaw)
				if taskNorm == "" {
					fmt.Println("Empty spaces aren't allowed. Please enter a valid task.")
					continue
				}
				insertTask(taskNorm)
				fmt.Println("Task added!")
				break
			}
		case 2:
			if len(todos) == 0 {
				fmt.Println("No tasks available to remove.")
				break
			}
			displayTasks()
			fmt.Print("Enter task number to remove: ")
			numRaw, _ := reader.ReadString('\n')
			numRaw = strings.TrimSpace(numRaw)
			index, err := strconv.Atoi(numRaw)
			if err != nil || index < 1 || index > len(todos) {
				fmt.Println("Invalid task number!")
				break
			}
			deleteTask(index - 1)
			fmt.Println("Task removed.")
		case 3:
			fmt.Println("Task List:")
			displayTasks()
		case 4:
			if len(todos) == 0 {
				fmt.Println("No tasks available to toggle.")
				break
			}
			displayTasks()
			fmt.Print("Enter task number to toggle status: ")
			numRaw, _ := reader.ReadString('\n')
			numRaw = strings.TrimSpace(numRaw)
			index, err := strconv.Atoi(numRaw)
			if err != nil || index < 1 || index > len(todos) {
				fmt.Println("Invalid task number!")
				break
			}
			toggleTaskStatus(index - 1)
			fmt.Println("Task status toggled.")
		case 0:
			saveTaskData()
			fmt.Println("Exiting. Bye!")
			return
		default:
			fmt.Println("Unknown option. Please try again.")
		}
		saveTaskData()
	}
}
