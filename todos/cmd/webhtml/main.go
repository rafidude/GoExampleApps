package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Todo struct {
	Title string
	Done  bool
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", home)

	app.Listen(":3000")
}

func home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Hello, World!",
		"Description": "This is a template inside layout.",
		"Todos": []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	})
}
