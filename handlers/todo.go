package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karledenstal/golang-todo/config"
	"github.com/karledenstal/golang-todo/entities"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []entities.Todo

	config.Database.Find(&todos)
	return c.Status(200).JSON(todos)
}

func GetDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo entities.Todo

	result := config.Database.Find(&todo, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(todo)
}
