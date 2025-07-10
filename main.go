package main

import (
	"Task_manager/config"
	"Task_manager/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	routes.TaskRoutes(app)

	app.Listen(":3000")
}
