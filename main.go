package main

import (
	"github.com/tot0p/Hackaton-25-Back/internal/api"

	"github.com/tot0p/env"
)

func init() {
	err := env.Load()
	if err != nil {
		return
	}
}

func main() {
	port := ":" + env.Get("PORT")
	if port == ":" {
		port = ":8080"
	}
	filename := env.Get("DB_FILE_PATH")
	if filename == "" {
		filename = "database.db"
	}
	url := env.Get("URL_OF_API_CO2")
	if url == "" {
		url = "http://localhost:8000"
	}
	Client := api.InitApi(port, filename)
	Client.RegisterHandlers(url)
	Client.Start()
}
