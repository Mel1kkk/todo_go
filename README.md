# Todo List (Terminal App)

## Git Repo
https://github.com/Mel1kkk/todo_go

## Information
- **Student Name**: *Miras Berikov*
- **Student ID**: *220103187*
- **Course Name**: *INF 368 - Go Lang*
- **Practice Group**: *05-P*

## About
This is a simple terminal Todo List app written in Go.

The app allows you to:
- Add new tasks.
- List all tasks.
- Mark tasks as done or undone.
- Remove tasks.

The tasks are stored in a local JSON file (`tasks.json`).  
When the app starts, it loads the tasks from the file.  
When you add, delete or toggle tasks, the file is updated.

## Features
- Spaces between words are fixed. (Example: `I     go    home` → `I go home`)
- Empty tasks (only spaces) are not allowed.
- You can toggle task status easily (done ↔ undone).
- JSON storage keeps your tasks safe between runs.

## How to run
You need to run both Go files together:

```terminal
go run main.go todo.go
