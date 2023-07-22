package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lasa.ai/todoCRUD/database"
)

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Hello, HTML!",
		"Description": "Template with No Layout.",
		"Todos": []database.Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	})
}

func Todo(c *fiber.Ctx) error {
	arrTodos, err := database.GetTodos()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Render("todo", fiber.Map{
		"Todos": arrTodos,
	}, "layouts/main")
}

func SaveTodo(c *fiber.Ctx) error {
	todo := c.FormValue("title")
	completed := c.FormValue("completed")

	done := false
	if completed == "true" {
		done = true
	}

	id, err := database.SaveTodo(todo, done)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"id": id, "title": todo, "completed": completed,
		"message": "Todo is successfully saved",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := c.FormValue("title")
	completed := c.FormValue("completed")
	done := false
	if completed == "true" {
		done = true
	}
	err := database.UpdateTodo(id, todo, done)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"id": id, "title": todo, "completed": completed,
		"message": "Todo is successfully updated",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := database.DeleteTodo(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Todo is successfully deleted",
	})
}
