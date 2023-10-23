package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"trip-reservation/db"
	"trip-reservation/models"
)

type UserHandler struct {
	userStore db.UserStore
}

// NewUserHandler is a "User Factory"
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (uh *UserHandler) GetUserHandler(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)

	user, err := uh.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (uh *UserHandler) GetUsersHandler(c *fiber.Ctx) error {
	test_u := models.User{
		ID:    "10",
		FName: "Mathew",
		LName: "Bravo",
	}
	return c.JSON(test_u)
}
