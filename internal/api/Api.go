package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/tot0p/Hackaton-25-Back/internal/db"
	"strconv"
	"time"
)

type Api struct {
	db   *db.Database
	app  *fiber.App
	port string
}

func (api *Api) Start() {
	defer api.db.Close()
	err := api.app.Listen(api.port)
	if err != nil {
		return
	}
}

func InitApi(port, filename string) *Api {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Local",
	}))

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

	return &Api{db: db.CreateDatabase(filename), app: app, port: port}
}
