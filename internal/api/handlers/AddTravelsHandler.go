package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/models/APIInput"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

// AddTravelsHandler returns a handler function that adds a travel to the database
func AddTravelsHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := utils.LoadUserJWT(c)
		var Input APIInput.Travel
		err := c.BodyParser(&Input)
		if err != nil {
			log.Errorw("Error parsing body", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
		}

		travel := DBModels.Travel{
			UserUUID:      user.UUID,
			StartLocation: Input.StartLocation,
			EndLocation:   Input.EndLocation,
			Distance:      Input.Distance,
			Duration:      Input.Duration,
			CO2:           Input.CO2,
			TransportType: Input.TransportType,
		}
		err = db.AddTravel(travel)
		if err != nil {
			log.Errorw("Error adding travel", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Travel added"})
	}
}
