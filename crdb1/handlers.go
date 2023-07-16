package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type user struct {
	Name string
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var users []string
	rows, err := db.Query("SELECT * FROM names")
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&res)
		users = append(users, res)
	}
	return c.Render("index", fiber.Map{
		"Users": users,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newUser := user{}
	if err := c.BodyParser(&newUser); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	if newUser.Name != "" {
		uuid, _ := uuid.NewRandom()
		_, err := db.Exec("INSERT into names VALUES ($1, $2)", uuid, newUser.Name)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}

	return c.Redirect("/")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	oldName := c.Params("name")
	newName := user{}

	if err := c.BodyParser(&newName); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	db.Exec("UPDATE names SET Name=$1 WHERE Name=$2", newName.Name, oldName)
	return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	userToDelete := c.Params("name")

	db.Exec("DELETE from names WHERE Name=$1", userToDelete)
	return c.SendString("deleted")
}
