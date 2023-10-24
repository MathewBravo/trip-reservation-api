package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"trip-reservation/api"
	"trip-reservation/db"
	"trip-reservation/helpers"
)

const uri = "mongodb://localhost:27017"
const DB_NAME = "trip-reservation"
const USER_COLLECTION = "users"

func main() {

	listenAddr := flag.String("listenAddr", ":8081", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(helpers.CONFIG)
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", userHandler.GetUsersHandler)
	apiv1.Get("/user/:id", userHandler.GetUserHandler)

	apiv1.Post("/user", userHandler.CreateUserHandler)

	err = app.Listen(*listenAddr)
	if err != nil {
		log.Fatal("Could not listen on port :8081")
	}
}
