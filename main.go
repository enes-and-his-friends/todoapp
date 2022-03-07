package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://db:27017/"

func main() {

	fmt.Println("db")

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	/*
		app.Get("/task", func(c *fiber.Ctx) error {
			return c.SendString("Task created!")
		})
	*/
	mongo_bson_D := bson.D{
		bson.E{Key: "id", Value: 1},
		bson.E{Key: "username", Value: "enestuzlu"},
		bson.E{Key: "password", Value: "123456"},
		bson.E{Key: "todos", Value: bson.D{
			bson.E{Key: "task_id", Value: 1},
			bson.E{Key: "task_name", Value: "Kitap okunacak"},
			bson.E{Key: "done", Value: false},
		},
		},
	}

	app.Get("/task", func(ctx *fiber.Ctx) error {
		quickstartDatabase := client.Database("quickstart")
		usersCollection := quickstartDatabase.Collection("users")
		go_ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		_, err := usersCollection.InsertOne(go_ctx, mongo_bson_D)
		if err != nil {
			fmt.Println("Task cannot be created due to error: ", err)
			return ctx.SendString("Task cannot be created due to error !")
		}
		return ctx.SendString("Task added !")
	})

	app.Listen(":3000")
}
