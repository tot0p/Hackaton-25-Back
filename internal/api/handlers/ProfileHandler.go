package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

func ProfileHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// load user from token
		user := utils.LoadUserJWT(c)

		//todo calulate user stats

		return c.JSON(fiber.Map{"username": user.Username})
	}
}
