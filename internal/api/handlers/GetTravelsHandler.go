package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

// GetTravelsHandler returns a handler function that returns the travels of the user
func GetTravelsHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := utils.LoadUserJWT(c)
		travels, err := db.GetTravels(user.UUID)
		if err != nil {
			log.Errorw("Error getting travels", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		return c.JSON(travels)
	}
}
