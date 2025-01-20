package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
)

func LoadUserJWT(c *fiber.Ctx) DBModels.User {
	// handle error
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var User DBModels.User
	User.UUID = claims["UserUUID"].(string)
	User.Username = claims["Username"].(string)
	return User
}
