package database

import (
	"gorm.io/gorm"
)

type Database struct {
	instance DBInterface
}

func New(db DBInterface) Database {
	return Database{
		instance: db,
	}
}

func (database Database) Connect() (*gorm.DB, error) {
	return database.instance.Connect()
}

func (database Database) Migrate(models map[string]interface{}) error {
	return database.instance.Migrate(models)
}
