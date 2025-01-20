package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

func GetTravelsHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := utils.LoadUserJWT(c)
		travels, err := db.GetTravels(user.UUID)
		if err != nil {
			return err
		}
		return c.JSON(travels)
	}
}
