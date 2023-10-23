package api

import (
	"github.com/gofiber/fiber/v2"
	"trip-reservation/models"
)

func GetUsersHandler(c *fiber.Ctx) error {
	test_u := models.User{
		ID:    "10",
		FName: "Mathew",
		LName: "Bravo",
	}
	return c.JSON(test_u)
}

func GetUserHandler(c *fiber.Ctx) error {
	return c.JSON("Mathew: 120")
}
