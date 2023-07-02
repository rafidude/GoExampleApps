package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render a template named 'index.html' with content
		return c.Render("index", fiber.Map{
			"Title":       "Hello, World!",
			"Description": "This is a template.",
		})
	})
	// Start server on port 5000
	app.Listen(":5000")
}
