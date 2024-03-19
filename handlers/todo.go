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

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo entities.Todo

	result := config.Database.Find(&todo, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(todo)
}

func AddTodo(c *fiber.Ctx) error {
	todo := new(entities.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&todo)
	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(entities.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&todo)

	return c.Status(200).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	config.Database.Delete(&entities.Todo{}, id)
	return c.SendStatus(200)
}
