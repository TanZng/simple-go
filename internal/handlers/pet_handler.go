package grpc

import (
	"context"
	"errors"
	core_pet "simple-go/internal/core/pet"
	log "simple-go/logger"
	"simple-go/validation"

	"github.com/google/uuid"
)

type PetHandler struct {
	PetService core_pet.PetServiceInterface
	logger     log.Logger
}

func NewPetHandler(petService core_pet.PetServiceInterface) PetHandler {
	return PetHandler{
		PetService: petService,
		logger:     log.NewLogger(),
	}
}

func (h PetHandler) GetPet(ctx context.Context, givenId string) (*core_pet.Pet, error) {

	petUUID, err := uuid.Parse(givenId)
	h.logger.Infow("[userHandler.GetUser] Incoming GetPet request", "petUUID", petUUID, "given", givenId)

	if err != nil {
		h.logger.Error("[userHandler.GetUser] Error while parsing req id", err)
		err := errors.New(validation.INVALID_UUID_MESSAGE)
		return &core_pet.Pet{}, err
	}

	foundPet, err := h.PetService.GetByID(ctx, petUUID)

	if err != nil {
		h.logger.Error("[petHandler.GetPet] Error while fetching pet on pet service GetById", err)
		return &core_pet.Pet{}, err
	}

	return foundPet, nil
}

func (h PetHandler) AddPet(ctx context.Context, givenPet core_pet.Pet) (*core_pet.Pet, error) {
	h.logger.Infow("[petHandler.AddPet] Incoming AddPet request", "pet", givenPet)

	if _, err := h.PetService.Create(ctx, &givenPet); err != nil {
		h.logger.Error("[petHandler.AddPet] Error while creating pet on pet service Create", err)
		return &core_pet.Pet{}, err
	}

	h.logger.Infow("[petHandler.AddPet] Created pet", "pet", givenPet)

	return &givenPet, nil
}
