package main

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type todoItem struct {
	Description string
	Completed   bool
}

var todos []todoItem
var fileName = "tasks.json"

func loadTaskData() {
	f, err := os.Open(fileName)
	if err != nil {
		todos = []todoItem{}
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		todos = []todoItem{}
		return
	}
	json.Unmarshal(data, &todos)
}

func saveTaskData() {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(fileName, data, 0644)
}

func insertTask(desc string) {
	newTask := todoItem{Description: desc, Completed: false}
	todos = append([]todoItem{newTask}, todos...)
}

func deleteTask(index int) {
	if index < 0 || index >= len(todos) {
		return
	}
	todos = append(todos[:index], todos[index+1:]...)
}

func toggleTaskStatus(index int) {
	if index < 0 || index >= len(todos) {
		return
	}
	todos[index].Completed = !todos[index].Completed
}

func displayTasks() {
	for i, task := range todos {
		checkMark := "X"
		if task.Completed {
			checkMark = "✓"
		}
		println(i+1, ". [", checkMark, "]", task.Description)
	}
}

func normalizeSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}
