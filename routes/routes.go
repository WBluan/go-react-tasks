package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/my-tasks/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/todos", controllers.GetTodos)
	app.Post("/api/todos", controllers.CreateTodos)
	app.Patch("/api/todos/:id", controllers.UpdateTodoByID)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)
}
