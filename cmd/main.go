package main

import (
	"fmt"
	"samplego/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("in main fn() before DB connection ..")
	database.ConnectDb()
	fmt.Println("in main fn() after DB connection ..")

	app := fiber.New()

	setupRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	fmt.Println("in main fn()..")
	// 	return c.SendString("Sample GO App from main!")
	// })

	app.Listen(":3000")
}
