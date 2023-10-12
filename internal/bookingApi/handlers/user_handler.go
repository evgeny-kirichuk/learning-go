package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Evgeny Kirichuk"})
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Evgeny Kirichuk"})
}