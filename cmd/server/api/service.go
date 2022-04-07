package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	AuthorizeToken(token string) error
}

type ApiService struct {
	api  *fiber.App
	auth AuthService
}

func New(api *fiber.App) ApiService {
	return ApiService{
		api: api,
	}
}

func (s ApiService) withAuthorization(c *fiber.Ctx, next func(c *fiber.Ctx) error) error {
	token := c.Get("Authorization") // Retrieves the authorization header
	if err := s.auth.AuthorizeToken(token); err != nil {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}
	return next(c)
}

func (s ApiService) Listen(addr string) error {
	return s.api.Listen(addr)
}

func (s ApiService) GetPublic(route string, handler func(*fiber.Ctx) error) {
	s.api.Get(route, func(c *fiber.Ctx) error { return handler(c) })
}

func (s ApiService) PostPublic(route string, handler func(*fiber.Ctx) error) {
	s.api.Post(route, func(c *fiber.Ctx) error { return handler(c) })
}
