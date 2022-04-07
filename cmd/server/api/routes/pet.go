package routes

import (
	"context"
	"net/http"
	"simple-go/cmd/server/api"
	core_pet "simple-go/internal/core/pet"
	handler "simple-go/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

type petRoutes struct{}

func PetRoutes() petRoutes {
	return petRoutes{}
}

func (u petRoutes) getPet(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	myHandler := handler.NewHandlers()
	petController := myHandler.PetHandler()
	givenContext := context.Background()
	pet, _ := petController.GetPet(givenContext, id)

	return ctx.JSON(pet)
}

func (u petRoutes) postPet(ctx *fiber.Ctx) error {
	pet, err := u.parseBody(ctx)
	if err != nil {
		return err
	}

	errors := api.ValidateStruct(*pet)
	if errors != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errors)
	}

	myHandler := handler.NewHandlers()
	petController := myHandler.PetHandler()
	givenContext := context.Background()
	pet, _ = petController.AddPet(givenContext, *pet)

	return ctx.Status(http.StatusCreated).JSON(pet)
}

func (u petRoutes) parseBody(ctx *fiber.Ctx) (*core_pet.Pet, error) {
	data := new(core_pet.Pet)
	err := api.BodyParser(ctx, &data)
	return data, err
}

func (u petRoutes) RegisterRoutes(api *api.ApiService) {
	api.GetPublic("/pet/:id", u.getPet)
	api.PostPublic("/pet", u.postPet)
}
