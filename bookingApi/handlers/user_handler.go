package handlers

import (
	"context"
	"learning_go/bookingApi/db"
	"learning_go/bookingApi/types"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUserById(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
		ctx = context.Background()
	)

	user,err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Evgeny",
		LastName:  "Kirichuk",
		Age:       30,
	}
	return c.JSON(u)
}

