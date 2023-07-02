package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"lasa.ai/fiber1/book"
)

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, Fiber12!")
}

func setupRoutes(app *fiber.App) {
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, Fiber12!")
	// })
	app.Get("/", home)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book", book.GetBooks)

}

func main() {
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
