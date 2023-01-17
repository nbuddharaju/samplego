package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nbuddharaju/samplego/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sample GO App!")
	})

	app.Listen(":3000")
}
