package handlers

import (
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"github.com/tot0p/Hackaton-25-Back/internal/models/APIInput"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
	"time"
)

func CreateUserHandler(db *DBManager.DBManager, cert *rsa.PrivateKey) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//Parse body into struct
		var Input APIInput.CreateUser
		err := c.BodyParser(&Input)
		if err != nil {
			log.Errorw("Error parsing body", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
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
			log.Errorw("Error hashing password", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		// Create user
		err = db.CreateUser(Input.Username, Input.Password)
		if err != nil {
			log.Errorw("Error creating user", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}

		// Get user by UUID
		user, err := db.GetUserByUsername(Input.Username)
		if err != nil {
			log.Errorw("Error getting user", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		exp := utils.GetExp()
		token, err := utils.CreateTokenJWT(*user, cert, exp)
		if err != nil {
			log.Errorw("Error creating token", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token, "exp": exp - time.Now().Unix()})
	}
}
