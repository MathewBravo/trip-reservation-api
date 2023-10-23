package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"trip-reservation/api"
	"trip-reservation/models"
)

const uri = "mongodb://localhost:27017"
const DB_NAME = "trip-reservation"
const USER_COLLECTION = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	col := client.Database(DB_NAME).Collection(USER_COLLECTION)

	u := models.User{
		FName: "Mathew",
		LName: "Bravo",
	}
	res, err := col.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	var testUser models.User

	err = col.FindOne(ctx, bson.M{}).Decode(&testUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(testUser)

	listenAddr := flag.String("listenAddr", ":8081", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", api.GetUsersHandler)
	apiv1.Get("/user/:id", api.GetUserHandler)

	err = app.Listen(*listenAddr)
	if err != nil {
		log.Fatal("Could not listen on port :8081")
	}
}
