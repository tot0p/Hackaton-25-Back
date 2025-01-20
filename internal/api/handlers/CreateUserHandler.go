package handlers

import (
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/models/APIInput"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

func CreateUserHandler(db *DBManager.DBManager, cert *rsa.PrivateKey) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//Parse body into struct
		var Input APIInput.CreateUser
		err := c.BodyParser(&Input)
		if err != nil {
			return err
		}

		// Check if Body is valid
		if Input.Username == "" || Input.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
		}
		// Check if user already exist
		ok, err := db.CheckIfUserExist(Input.Username)
		if err != nil {
			return err
		} else if ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exist"})
		}

		// hash password
		Input.Password, err = utils.HashPassword(Input.Password)
		if err != nil {
			return err
		}
		// Create user
		err = db.CreateUser(Input.Username, Input.Password)
		if err != nil {
			return err
		}

		// Get user by UUID
		user, err := db.GetUserByUsername(Input.Username)
		if err != nil {
			return err
		}

		token, err := utils.CreateTokenJWT(*user, cert)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
	}
}
