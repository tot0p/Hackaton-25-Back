package api

import (
	"github.com/tot0p/Hackaton-25-Back/internal/api/handlers"
)

func (api *Api) RegisterHandlers() {
	api.app.Get("/ping", handlers.PingHandler())
	api.app.Post("/user", handlers.CreateUserHandler(api.db))
}
