package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetCalculC02Handler(urlAPi string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := c.Body()
		apiURL := urlAPi + "/api/footprint"
		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// unmarshal the response
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return err
		}
		newStr := buf.String()
		var data map[string]interface{}
		err = json.Unmarshal([]byte(newStr), &data)
		if err != nil {
			return err
		}
		return c.JSON(data)

	}
}
