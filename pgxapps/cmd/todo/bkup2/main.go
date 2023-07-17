package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

// add a struct to hold the task data
type Task struct {
	id          string
	description string
}

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
}

func update_task() {
	// update a single task
	uid, err := uuid.Parse("999ebafd-fd1c-41ee-9466-e88b8496fec2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to convert task_num into uuid: %v\n", err)
		os.Exit(1)
	}

	description := "Learn Kindness"
	fmt.Println("Updating task", uid, "to", description)
	_, err = db.Exec("update tasks set description = $1 where id = $2", description, uid)
	// if err return the following error
	if err != nil {
		fmt.Printf("unable to update task: %v\n", err)
	}
	fmt.Println("Task updated")
}

func add_task() {
	description := "Learn Manners"
	_, err := db.Exec("insert into tasks(description) values($1)", description)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Task added")
}

func list_tasks1() {
	// select all rows from tasks table and print description
	rows, err := db.Query("select * from tasks")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to query tasks: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var description string
		var id string
		err = rows.Scan(&id, &description)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan row: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s. %s\n", id, description)
	}
}

func list_tasks2() {
	// select rows from tasks and populate a slice of Task structs
	rows, err := db.Query("select * from tasks")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to query tasks: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.id, &task.description)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan row: %v\n", err)
			os.Exit(1)
		}
		tasks = append(tasks, task)
	}
	// print the slice of Task structs
	for _, task := range tasks {
		fmt.Printf("%s. %s\n", task.id, task.description)
	}
}
