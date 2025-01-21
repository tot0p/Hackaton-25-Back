package utils

import (
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
	"time"
)

// LoadUserJWT loads the user from the JWT token
func LoadUserJWT(c *fiber.Ctx) DBModels.User {
	// handle error
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var User DBModels.User
	User.UUID = claims["UserUUID"].(string)
	User.Username = claims["Username"].(string)
	return User
}

// CreateTokenJWT creates a JWT token
func CreateTokenJWT(user DBModels.User, cert *rsa.PrivateKey, exp int64) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"UserUUID": user.UUID,
		"Username": user.Username,
		"exp":      exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(cert)

}

// GetExp returns the expiration time of the token
func GetExp() int64 {
	return time.Now().Add(time.Hour * 72).Unix()
}
