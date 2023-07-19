package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"lasa.ai/todoCRUD/routes"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", routes.Home)
	app.Get("/todos", routes.Todo)
	app.Post("/todos", routes.SaveTodo)
	app.Post("/todos/:id", routes.UpdateTodo)
	app.Delete("/todos/:id", routes.DeleteTodo)

	app.Listen(":3000")
}
