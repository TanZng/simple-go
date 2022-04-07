package database

import "gorm.io/gorm"

type DBInterface interface {
	Migrate(models map[string]interface{}) error
	Connect() (*gorm.DB, error)
}

type DatabaseInstance struct {
	PostgreSQL DBInterface
}

type Store struct {
	SqlDB DatabaseInstance
}

func NewStore() Store {
	return Store{
		SqlDB: DatabaseInstance{
			PostgreSQL: New(NewPostgreInstance()),
		},
	}
}
