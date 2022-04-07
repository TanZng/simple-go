package database

import (
	"context"
	"simple-go/internal/core/pet"
)

type DBSeedInterface interface {
	Migrate(models map[string]interface{}) error
}

type DBSeed struct {
	db DBSeedInterface
}

func NewSeed(database DBSeedInterface) DBSeed {
	return DBSeed{
		db: database,
	}
}

func (s DBSeed) Setup(ctx context.Context) {
	err := s.db.Migrate(
		s.models(),
	)

	if err != nil {
		panic(err)
	}
}

func (s DBSeed) models() map[string]interface{} {
	return map[string]interface{}{
		"pet.Pet": pet.Pet{},
	}
}
