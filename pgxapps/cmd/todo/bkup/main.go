package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

/*
CREATE TABLE tasks (

	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	description TEXT NOT NULL

);
*/

func main() {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "list":
		err = listTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list tasks: %v\n", err)
			os.Exit(1)
		}

	case "add":
		err = addTask(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err)
			os.Exit(1)
		}

	case "update":
		n, err := uuid.Parse(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into uuid: %v\n", err)
			os.Exit(1)
		}
		err = updateTask(n, os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to update task: %v\n", err)
			os.Exit(1)
		}

	case "remove":
		n, err := uuid.Parse(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into uuid: %v\n", err)
			os.Exit(1)
		}
		err = removeTask(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to remove task: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		printHelp()
		os.Exit(1)
	}
}

func listTasks() error {
	rows, _ := conn.Query(context.Background(), "select * from tasks")

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
	_, err := conn.Exec(context.Background(), "insert into tasks(description) values($1)", description)
	return err
}

func updateTask(itemNum uuid.UUID, description string) error {
	_, err := conn.Exec(context.Background(), "update tasks set description=$1 where id=$2", description, itemNum)
	return err
}

func removeTask(itemNum uuid.UUID) error {
	_, err := conn.Exec(context.Background(), "delete from tasks where id=$1", itemNum)
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
