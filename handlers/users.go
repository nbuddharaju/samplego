package handlers

import "github.com/gofiber/fiber/v2"

func GET(c *fiber.Ctx) error {
	return c.SendString("Sample App!")
}

func PUT(c *fiber.Ctx) error {
	return c.SendString("Sample App PUT!")
}

func POST(c *fiber.Ctx) error {
	return c.SendString("Sample App POST!")
}

func DELETE(c *fiber.Ctx) error {
	return c.SendString("Sample App DELETE!")
}
