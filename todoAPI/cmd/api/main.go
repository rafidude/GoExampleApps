package main

import (
	"todoAPI/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Todo handlers
	app.Get("/todos", handlers.GetTodos)
	app.Post("/todos", handlers.SaveTodo)
	app.Put("/todos", handlers.UpdateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)

	// Some other entity handlers
	app.Listen(":3000")
}
