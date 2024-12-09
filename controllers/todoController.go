package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/my-tasks/models"
	"github.com/wbluan/my-tasks/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(c *fiber.Ctx) error {
	todos, err := repository.GetTodos()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch todos"})
	}
	return c.JSON(todos)
}

func CreateTodos(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	newTodo, err := repository.CreateTodos(todo)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(newTodo)
}

func UpdateTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo's ID"})
	}

	err = repository.UpdateTodoByID(objectID)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo's ID"})
	}
	err = repository.DeleteTodoByID(objectID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
