package grpc

import (
	db "simple-go/database"
	"simple-go/internal/core/pet"
	log "simple-go/logger"
)

type Handlers struct {
	store  db.Store
	logger log.Logger
}

func NewHandlers() Handlers {
	return Handlers{
		store:  db.NewStore(),
		logger: log.NewLogger(),
	}
}

func (h Handlers) PetHandler() PetHandler {
	h.logger.Info("[PetHandler] Starting")
	repository, err := pet.NewPetRepository(h.store.SqlDB.PostgreSQL)

	if err != nil {
		h.logger.Error("[PetHandler] Error when creating PetRepository", err)
		panic(err)
	}

	petHandler := NewPetHandler(pet.NewPetService(repository))

	return petHandler

}
