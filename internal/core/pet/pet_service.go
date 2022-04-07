package pet

import (
	"context"
	"errors"
	"simple-go/logger"
	"simple-go/validation"

	"github.com/google/uuid"
)

type PetServiceInterface interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Pet, error)
	Create(ctx context.Context, pet *Pet) (uuid.UUID, error)
}

// PetServiceT es el Servicio para hacer operaciones sobre el Pet.
type PetServiceT struct {
	store  PetStoreInterface
	logger logger.Logger
}

func NewPetService(repository PetStoreInterface) PetServiceT {
	return PetServiceT{
		store:  repository,
		logger: logger.NewLogger(),
	}
}

func (bs PetServiceT) GetByID(ctx context.Context, id uuid.UUID) (*Pet, error) {
	pet, err := bs.store.GetByID(id)
	if err != nil {
		err = errors.New(validation.RESOURCE_NOT_FOUND_WITH_ID)
		bs.logger.Error("[petService.GetByID] Error fetching pet by id", err)
		return nil, err
	}

	bs.logger.Infow("[petService.GetByID] User found", "pet", pet)

	return pet, nil
}

func (bs PetServiceT) Create(ctx context.Context, pet *Pet) (uuid.UUID, error) {
	newID, err := bs.store.Create(pet)

	if err != nil {
		bs.logger.Error("[petService.Create] Error creating pet", err)
		return uuid.UUID{}, err
	}

	if newID.String() == validation.EMPTY_UUID {
		err = errors.New(validation.RESOURCE_CREATED_WITHOUT_UUID)
		bs.logger.Error("[petService.Create] Created pet has no uuid", err)
		return uuid.UUID{}, err
	}

	pet.ID = newID
	bs.logger.Infow("[petService.Create] Created pet", "pet", pet)

	return newID, nil
}
