package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func getDB() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	db, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*
CREATE TABLE tasks (

	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	description TEXT NOT NULL

);
*/

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	var err error

	switch os.Args[1] {
	case "list":
		err = listTasks()
	case "add":
		err = addTask(os.Args[2])
	case "update":
		n, err := uuid.Parse(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into uuid: %v\n", err)
			os.Exit(1)
		}
		err = updateTask(n, os.Args[3])
	case "remove":
		n, err := uuid.Parse(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into uuid: %v\n", err)
			os.Exit(1)
		}
		err = removeTask(n)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		printHelp()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
		os.Exit(1)
	}
}

func listTasks() error {
	db, err := getDB()
	if err != nil {
		return err
	}

	rows, err := db.Query("select * from tasks")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		fmt.Printf("%s. %s\n", id, description)
	}

	return rows.Err()
}

func addTask(description string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("insert into tasks(description) values($1)", description)
	return err
}

func updateTask(itemNum uuid.UUID, description string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("update tasks set description=$1 where id=$2", description, itemNum)
	return err
}

func removeTask(itemNum uuid.UUID) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("delete from tasks where id=$1", itemNum)
	return err
}

func printHelp() {
	fmt.Print(`Todo pgx demo

Usage:

  todo list
  todo add task
  todo update task_num item
  todo remove task_num

Example:

  todo add 'Learn Go'
  todo list
`)
}
