package handlers

import "github.com/gofiber/fiber/v2"

// Ping returns a pong message
func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// PingHandler returns a handler function that returns a pong message
func PingHandler() func(c *fiber.Ctx) error {
	return Ping
}
