package api

import (
	"crypto/rand"
	"crypto/rsa"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"time"
)

type Api struct {
	db     *DBManager.DBManager
	app    *fiber.App
	port   string
	cert   *rsa.PrivateKey
	secure fiber.Handler
}

func (api *Api) Start() {
	api.db.Open()
	api.db.Init()
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
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
	rng := rand.Reader
	var err error
	cert, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
	secure := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    cert.Public(),
		}})
	return &Api{db: DBManager.NewDBManager(filename), app: app, port: port, cert: cert, secure: secure}
}
