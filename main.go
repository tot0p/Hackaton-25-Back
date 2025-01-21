package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/tot0p/Hackaton-25-Back/internal/api"
	"os"

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
		log.Fatal("URL_OF_API_CO2 is not set")
		os.Exit(1)
	}
	Client := api.InitApi(port, filename)
	Client.RegisterHandlers(url)
	Client.Start()
}
