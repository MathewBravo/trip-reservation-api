package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8081", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiv1 := app.Group("api/v1")

	app.Get("/foo", handleFoo)
	apiv1.Get("/user", handleUser)

	err := app.Listen(*listenAddr)
	if err != nil {
		log.Fatal("Could not listen on port :8081")
	}
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "app is working"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Mathew Bravo"})
}
