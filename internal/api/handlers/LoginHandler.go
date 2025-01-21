package handlers

import "C"
import (
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/models/APIInput"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
	"time"
)

// LoginHandler returns a handler function that logs in the user
func LoginHandler(db *DBManager.DBManager, cert *rsa.PrivateKey) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var Input APIInput.Login
		err := c.BodyParser(&Input)
		if err != nil {
			log.Errorw("Error parsing body", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
		}
		if Input.Username == "" || Input.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
		}
		user, err := db.GetUserByUsernameWithPass(Input.Username)
		if err != nil {
			log.Errorw("Error getting user", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		} else if user == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
		}

		if !utils.CheckPasswordHash(Input.Password, user.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
		}

		exp := utils.GetExp()
		token, err := utils.CreateTokenJWT(*user, cert, exp)
		if err != nil {
			log.Errorw("Error creating token", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp - time.Now().Unix()})
	}
}
