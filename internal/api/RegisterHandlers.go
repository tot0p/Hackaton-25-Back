package api

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/api/handlers"
)

func (api *Api) RegisterHandlers(urlApi string) {
	// public routes
	api.app.Get("/ping", handlers.PingHandler())
	api.app.Post("/user", handlers.CreateUserHandler(api.db, api.cert))
	api.app.Post("/login", handlers.LoginHandler(api.db, api.cert))

	// secure routes
	api.app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    api.cert.Public(),
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		},
	}))
	api.app.Get("/profile", handlers.ProfileHandler(api.db))
	api.app.Post("/travels", handlers.AddTravelsHandler(api.db))
	api.app.Get("/travels", handlers.GetTravelsHandler(api.db))
	api.app.Post("/CalculateCO2", handlers.GetCalculC02Handler(urlApi))

}
