//go:generate mockgen -destination=./pet_store_mocks_test.go -package=pet -source=./pet_store.go
package pet

import "github.com/google/uuid"

// PetStoreInterface es implementada por las clases PetServiceT y PetRepositoryT
type PetStoreInterface interface {
	GetByID(id uuid.UUID) (*Pet, error)
	Create(*Pet) (uuid.UUID, error)
}
