// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	// Get all records from postgreSQL
	app.Get("/employee", GetEmployee)

	// Add record into postgreSQL
	app.Post("/employee", PostEmployee)

	// Update record into postgreSQL
	app.Put("/employee", UpdateEmployee)

	// Delete record from postgreSQL
	app.Delete("/employee", DeleteEmployee)

	log.Fatal(app.Listen(":3000"))
}
