package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
)

func main() {
	// Create a new engine
	engine := pug.New("./views", ".pug")

	app := fiber.New(fiber.Config{
		//Views: html.New("./views", ".html"),
		Views: engine,
	})
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, WD!",
		}, "layout")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello, " + c.Params("name"))
	})

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
			"posts":   []string{"post 1", "post 2"},
		})
	})

	log.Fatal(app.Listen(":3000"))
}
