package handlers

import "github.com/gofiber/fiber/v2"

func ProfileHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Profile"})
	}
}
