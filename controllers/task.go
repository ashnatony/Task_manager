package controllers

import (
	"Task_manager/config"
	"Task_manager/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID := c.Locals("userId").(float64)
	task.UserID = uint(userID)

	if err := config.DB.Create(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create task"})
	}

	return c.JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	userID := c.Locals("userId").(float64)

	var tasks []models.Task
	if err := config.DB.Where("user_id = ?", uint(userID)).Find(&tasks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch tasks"})
	}

	config.DB.Where("user_id = ?", userID).Order(`
			CASE priority
				WHEN 'High' THEN 1
				WHEN 'Medium' THEN 2
				WHEN 'Low' THEN 3
				ELSE 4
			END,
			due_date ASC,
			completed ASC
		`).Find(&tasks)

	return c.JSON(tasks)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := uint(c.Locals("userId").(float64))

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	var input models.Task
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Completed = input.Completed
	task.Priority = input.Priority
	task.DueDate = input.DueDate

	config.DB.Save(&task)
	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := uint(c.Locals("userId").(float64))

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	config.DB.Delete(&task)
	return c.SendStatus(204)
}
