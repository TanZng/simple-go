package pet

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pet struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string         `json:"name" validate:"required"`
	Kind      string         `json:"kind" validate:"required"`
	CreatedAt time.Time      `json:"created-at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `json:"updated-at"`
	DeletedAt gorm.DeletedAt `json:"deleted-at" gorm:"index"`
}

func (pet *Pet) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	tx.Set("ID", &id)
	pet.ID = id
	return nil
}
