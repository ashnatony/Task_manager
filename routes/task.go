package routes

import (
	"Task_manager/controllers"
	"Task_manager/middleware"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App) {
	// Public routes for auth
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// Protected group for tasks
	api := app.Group("/api/tasks", middleware.Protected())
	api.Post("/", controllers.CreateTask)
	api.Get("/", controllers.GetTasks)
	api.Put("/:id", controllers.UpdateTask)
	api.Delete("/:id", controllers.DeleteTask)
}
