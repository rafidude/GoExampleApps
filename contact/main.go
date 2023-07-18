package main

import (
	"contact/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Post("/save", database.SaveContact)

	app.Listen(":3000")
}
