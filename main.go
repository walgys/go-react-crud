package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/walgys/go-react-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Use(cors.New())

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("gomongodb").Collection("users")

	app.Static("/", "./client/dist")
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		result, error := coll.Find(context.TODO(), bson.M{})
		if error != nil {
			panic(error)
		}
		for result.Next(context.TODO()) {
			var user models.User
			error := result.Decode(&user)
			if error != nil {
				panic(error)
			}
			users = append(users, user)
		}

		return c.JSON(fiber.Map{
			"users": users,
		})
	})
	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		c.BodyParser(&user)
		result, error := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})
		if error != nil {
			panic(error)
		}
		return c.JSON(fiber.Map{
			"result": result,
		})

	})

	app.Listen(":" + port)
}
