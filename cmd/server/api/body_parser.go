package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func BodyParser(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		c.Response().SetStatusCode(http.StatusBadRequest)
		return err
	}

	return nil
}
