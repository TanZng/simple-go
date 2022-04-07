//go:generate mockgen -destination=./pet_repository_mocks_test.go -package=pet -source=./pet_repository.go
package pet

import (
	"errors"
	"simple-go/logger"
	"simple-go/validation"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PetRepostitoryDBI es la interfaz que implementan las clases que
// implementan DBInterface (ej. PostgresRepo)
type PetRepostitoryDBI interface {
	Connect() (*gorm.DB, error)
}

type PetRepositoryT struct {
	db     gorm.DB
	logger logger.Logger
}

func NewPetRepository(db PetRepostitoryDBI) (PetRepositoryT, error) {
	sqlDB, err := db.Connect()

	if err != nil {
		return PetRepositoryT{}, err
	}

	return PetRepositoryT{
		db:     *sqlDB,
		logger: logger.NewLogger(),
	}, nil
}

func (br PetRepositoryT) GetByID(id uuid.UUID) (*Pet, error) {
	pet, err := br.findFromDatabase(id)

	if err != nil {
		br.logger.Error("[petRepository.GetByID] Error while fetching pet from database by id", err)
		return nil, err
	}

	br.logger.Infow("[petRepository.GetByID] pet from database", "pet", pet)

	return pet, nil
}

func (br PetRepositoryT) Create(pet *Pet) (uuid.UUID, error) {
	br.logger.Infow("[petRepository.Create] Inserting new pet", "pet", pet)
	tx := br.db.Begin()

	if err := tx.Create(pet).Error; err != nil {
		br.logger.Error("[petRepository.Create] Error while trying to add new pet to database", err)
		return uuid.UUID{}, err
	}
	err := tx.Commit().Error

	if err != nil {
		br.logger.Error("[petRepository.Create] Error in transaction commit", err)
		return uuid.UUID{}, err
	}

	br.logger.Infow("[petRepository.Create] Generated uuid for pet", "id", pet.ID)

	return pet.ID, nil
}

func (br PetRepositoryT) findFromDatabase(id uuid.UUID) (*Pet, error) {
	pet := Pet{}
	res := br.db.Find(&pet, id)

	if res.RowsAffected == 0 {
		err := errors.New(validation.RESOURCE_NOT_FOUND_WITH_ID)
		br.logger.Error("[petRepository.findFromDatabase] pet not found", err)
		return nil, err
	}

	if res.Error != nil {
		br.logger.Error("[petRepository.findFromDatabase] Error while finding pet from database", res.Error)
		return nil, res.Error
	}

	return &pet, nil
}
