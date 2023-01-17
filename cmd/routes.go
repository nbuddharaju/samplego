package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nbuddharaju/samplego/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
