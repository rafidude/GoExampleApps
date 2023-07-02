package main

import (
	"github.com/gofiber/fiber/v2"
	"lasa.ai/fiber2/api"
)

func main() {
	app := api.SetupRoute()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber API")
	})

	app.Listen(":5000")
}
