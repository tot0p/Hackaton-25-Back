package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
)

func ProfileHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Profile"})
	}
}
