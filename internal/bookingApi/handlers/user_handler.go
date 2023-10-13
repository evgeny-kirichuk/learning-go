package handlers

import (
	"learning_go/internal/bookingApi/types"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Evgeny",
		LastName:  "Kirichuk",
		Age:       30,
	}
	return c.JSON(u)
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Evgeny Kirichuk"})
}