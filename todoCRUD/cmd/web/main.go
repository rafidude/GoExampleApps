package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"lasa.ai/todoCRUD/handlers"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", handlers.Home)

	// Todo handlers
	app.Get("/todos", handlers.Todo)
	app.Post("/todos", handlers.SaveTodo)
	app.Post("/todos/:id", handlers.UpdateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)

	// Some other entity handlers
	app.Listen(":3000")
}
