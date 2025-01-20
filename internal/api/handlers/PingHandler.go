package handlers

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

func PingHandler() func(c *fiber.Ctx) error {
	return Ping
}
