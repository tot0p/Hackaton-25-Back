package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

// ProfileHandler returns a handler function that returns the profile of the user
func ProfileHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// load user from token
		user := utils.LoadUserJWT(c)
		CO2OfMonth, err := db.GetCO2OfMonth(user.UUID)
		if err != nil {
			log.Errorw("Error getting CO2 of month", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting CO2 of month"})
		}
		CO2OfPrevMonth, err := db.GetCO2OfPrevMonth(user.UUID)
		if err != nil {
			log.Errorw("Error getting CO2 of prev month", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting CO2 of prev month"})
		}
		if CO2OfPrevMonth == -1 {
			return c.JSON(fiber.Map{"username": user.Username, "CO2OfMonth": CO2OfMonth, "CO2OfPrevMonth": nil})
		}
		return c.JSON(fiber.Map{"username": user.Username, "CO2OfMonth": CO2OfMonth, "CO2OfPrevMonth": CO2OfPrevMonth})
	}
}
