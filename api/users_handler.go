package api

import (
	"fmt"
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

func (uh *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var params models.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	fmt.Println(params)
	if errs := params.Validate(); errs != nil {
		return c.JSON(errs)
	}
	user, err := models.CreateUserFromParams(params)
	if err != nil {
		return err
	}
	newUser, err := uh.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(newUser)
}

func (uh *UserHandler) GetUserHandler(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)

	user, err := uh.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (uh *UserHandler) GetUsersHandler(c *fiber.Ctx) error {
	users, err := uh.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
