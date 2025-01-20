package handlers

import "C"
import (
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/models/APIInput"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

func LoginHandler(db *DBManager.DBManager, cert *rsa.PrivateKey) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var Input APIInput.Login
		err := c.BodyParser(&Input)
		if err != nil {
			return err
		}
		if Input.Username == "" || Input.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
		}
		user, err := db.GetUserByUsernameWithPass(Input.Username)
		if err != nil {
			return err
		} else if user == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
		}

		if !utils.CheckPasswordHash(Input.Password, user.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
		}

		prevToken, err := db.JWTGetByUserUUIDIfNotExpired(user.UUID)
		if err != nil {
			return err
		}

		if prevToken != "" {
			return c.JSON(fiber.Map{"token": prevToken})
		}

		exp := utils.GetExp()
		token, err := utils.CreateTokenJWT(*user, cert, exp)
		if err != nil {
			return err
		}

		err = db.JWTRegister(user.UUID, token, exp)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp})
	}
}
