package handlers

import "C"
import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/models/APIInput"
)

func LoginHandler(db *DBManager.DBManager) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var Input APIInput.Login
		err := c.BodyParser(&Input)
		if err != nil {
			return err
		}
		if Input.Username == "" || Input.Password == "" || Input.Device == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
		}
		session, ok, err := db.Login(Input.Username, Input.Password, Input.Device)
		if err != nil {
			return err
		} else
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid password"})
		}
		
		return c.JSON(session)
	}
}
