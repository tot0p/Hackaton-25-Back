package api

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tot0p/Hackaton-25-Back/internal/DBManager"
	"time"
)

// Api is the main struct of the API
type Api struct {
	db   *DBManager.DBManager
	app  *fiber.App
	port string
	cert *rsa.PrivateKey
}

// Start starts the API
func (api *Api) Start() {
	api.db.Open()
	api.db.Init()
	defer api.db.Close()
	err := api.app.Listen(api.port)
	if err != nil {
		return
	}
}

// InitApi creates a new Api object
func InitApi(port, filename string) *Api {
	app := fiber.New()
	// add middleware to log requests
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Local",
	}))

	// add middleware to allow CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// add middleware to limit the number of requests
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
	// generate RSA key
	rng := rand.Reader
	var err error
	cert, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
	return &Api{db: DBManager.NewDBManager(filename), app: app, port: port, cert: cert}
}
