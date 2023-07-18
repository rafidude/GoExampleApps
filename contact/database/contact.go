package database

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func setupDatabase() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveContact(c *fiber.Ctx) error {
	db, err := setupDatabase()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer db.Close()

	email := c.FormValue("email")
	message := c.FormValue("message")

	_, err = db.Exec("INSERT INTO contacts (email, message) VALUES ($1, $2)", email, message)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Form submitted successfully!")
}
