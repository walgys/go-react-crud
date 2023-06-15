package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Use(cors.New())
	app.Static("/", "./client/dist")
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "user data"})
	})

	app.Listen(":" + port)
}
