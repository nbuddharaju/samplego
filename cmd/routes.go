package main

import (
	"fmt"
	"samplego/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	fmt.Println(" IN setup test routes.. ")
	app.Get("/", handlers.GET)
	app.Post("/", handlers.POST)
	app.Put("/", handlers.PUT)
	app.Delete("/", handlers.DELETE)
}
