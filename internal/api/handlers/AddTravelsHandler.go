package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/models/APIInput"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

func AddTravelsHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := utils.LoadUserJWT(c)
		var Input APIInput.Travel
		err := c.BodyParser(&Input)
		if err != nil {
			return err
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
			return err
		}

		return c.JSON(fiber.Map{"message": "Travel added"})
	}
}
