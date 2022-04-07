package server

import (
	"os"
	"simple-go/cmd/server/api"
	"simple-go/cmd/server/api/routes"

	"github.com/gofiber/fiber/v2"
)

func Init() {
	app := fiber.New()

	newApi := api.New(app)

	routes.New(&newApi, routes.PetRoutes()).Register()
	routes.New(&newApi, routes.HelloRoutes()).Register()

	newApi.Listen(os.Getenv("API_PORT"))
}
