package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/tot0p/Hackaton-25-Back/internal/handlers"
	"github.com/tot0p/env"
	"strconv"
	"time"
)

func init() {
	err := env.Load()
	if err != nil {
		return
	}
}

func main() {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        60,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).SendString("{\"error\": \"Too many requests\"}")
		},
	}))

	app.Use(cache.New(cache.Config{
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			newCacheTime, _ := strconv.Atoi(c.GetRespHeader("Cache-Time", "0"))
			return time.Second * time.Duration(newCacheTime)
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.OriginalURL())
		},
	}))

	// register routes
	app.Get("/ping", handlers.PingHandler())

	// start server
	port := ":" + env.Get("PORT")
	if port == "" {
		port = ":8080"
	}
	err := app.Listen(port)
	if err != nil {
		return
	}
}
