package handlers

import (
	"learning_go/bookingApi/db"
	"learning_go/bookingApi/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	)

	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, error := h.userStore.GetUsers(c.Context())
	if error != nil {
		return error
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var (
		values bson.M
		userID = c.Params("id")
	)
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&values); err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, values); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	newUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(newUser)
}
