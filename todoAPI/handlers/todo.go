package handlers

import (
	"todoAPI/database"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	arrTodos, err := database.GetTodos()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"Todos": arrTodos})
}

func SaveTodo(c *fiber.Ctx) error {
	todo := new(database.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id, err := database.SaveTodo(todo.Title, todo.Completed)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"id": id, "title": todo.Title, "completed": todo.Completed,
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	todo := new(database.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err := database.UpdateTodo(todo.Id, todo.Title, todo.Completed)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"id": todo.Id, "title": todo.Title, "completed": todo.Completed,
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
