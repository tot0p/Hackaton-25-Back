package api

import (
	"github.com/tot0p/Hackaton-25-Back/internal/api/handlers"
)

func (api *Api) RegisterHandlers() {
	// public routes
	api.app.Get("/ping", handlers.PingHandler())
	api.app.Post("/user", handlers.CreateUserHandler(api.db))
	api.app.Post("/login", handlers.LoginHandler(api.db, api.cert))

	// secure routes
	api.app.Use(api.secure)
	api.app.Get("/profile", handlers.ProfileHandler(api.db))

}
