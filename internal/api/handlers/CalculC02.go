package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
)

// GetCalculC02Handler returns a handler that sends a request to the API and returns the response
func GetCalculC02Handler(urlAPi string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := c.Body()
		apiURL := urlAPi + "/api/footprint"
		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Errorw("Error sending request to API", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			log.Errorw("Error reading response body", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		newStr := buf.String()
		var data map[string]interface{}
		err = json.Unmarshal([]byte(newStr), &data)
		if err != nil {
			log.Errorw("Error unmarshalling response body", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		err = resp.Body.Close()
		if err != nil {
			log.Errorw("Error closing response body", "error", err)
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		return c.JSON(data)
	}
}
