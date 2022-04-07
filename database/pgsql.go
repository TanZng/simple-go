package database

import (
	"fmt"
	"log"
	"os"
	"time"

	pfy_logger "simple-go/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

type PostgresRepo struct {
	logger pfy_logger.Logger
}

func NewPostgreInstance() PostgresRepo {
	return PostgresRepo{
		logger: pfy_logger.NewLogger(),
	}
}

func (pg PostgresRepo) Connect() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbName,
		dbPassword,
		dbSSLMode,
	)

	newGormLogger := gorm_logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gorm_logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gorm_logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: newGormLogger,
	})
	if err != nil {
		pg.logger.Error("[pgsql.Connect] Error while opening connection with pgSQL", err)
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		pg.logger.Error("[pgsql.Connect] Error while creating connection with pgSQL", err)
		return nil, err
	}

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(20)
	conn.SetConnMaxIdleTime(time.Hour)

	return db, nil
}

func (pg PostgresRepo) Migrate(models map[string]interface{}) error {
	conn, err := pg.Connect()
	if err != nil {
		return err
	}

	conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	for key, model := range models {
		pg.logger.Infow("[psql.Migrate] Migrating", "model", key)

		err = conn.Migrator().AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
