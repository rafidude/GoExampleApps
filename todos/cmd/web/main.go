package main

import (
	"lasa.ai/todos/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
)

func main() {
	engine := pug.New("./views", ".pug")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", home)
	app.Get("/todo", todo)
	app.Post("/save", database.SaveTodo)

	app.Listen(":3000")
}

func home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Hello, World!",
		"Description": "This is a template inside layout.",
	}, "layouts/main")
}

func todo(c *fiber.Ctx) error {
	return c.Render("todo", fiber.Map{
		"Title":       "Todo Items!",
		"Description": "Showing all todo items.",
	}, "layouts/main")
}
