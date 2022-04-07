package routes

import (
	"simple-go/cmd/server/api"

	"github.com/gofiber/fiber/v2"
)

type helloRoutes struct{}

func HelloRoutes() helloRoutes {
	return helloRoutes{}
}

func (u helloRoutes) getHelloWorld(ctx *fiber.Ctx) error {

	return ctx.JSON("Hello World")
}

func (u helloRoutes) RegisterRoutes(api *api.ApiService) {
	api.GetPublic("/hello-world", u.getHelloWorld)
}
