package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/Hackaton-25-Back/internal/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/ping", handlers.PingHandler())

	app.Listen(":8080")
}
