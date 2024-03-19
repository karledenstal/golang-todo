package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/karledenstal/golang-todo/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/todos", handlers.GetTodos)
	app.Get("/todos/:id", handlers.GetTodo)
	app.Post("/todos", handlers.AddTodo)
	app.Put("/todos/:id", handlers.UpdateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)

	log.Fatal(app.Listen(":9090"))
}
