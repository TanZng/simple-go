package test_helpers

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Seed struct {
	db   *gorm.DB
	path string
}

func NewSeed() *Seed {
	conn, _ := Connect()
	return &Seed{
		db:   conn,
		path: "../../",
	}
}

// func (s *Seed) Run() {
// 	s.createBeneficiarios()
// }

func (s *Seed) Beneficiarios() {
	s.createBeneficiosScenarios()
}

func (s *Seed) Clean() {
	s.truncateTables()
}

func (s *Seed) SetPath(path string) {
	s.path = path
}

func (s *Seed) executeSqlFromFile(filename string) {
	fd, _ := os.ReadFile(filename)
	fileContent := string(fd)
	s.db.Exec(fileContent)
}

func (s *Seed) createBeneficiosScenarios() {
	s.executeSqlFromFile(fmt.Sprintf("%stest_helpers/seed_scripts/beneficiario.sql", s.path))
}

func (s *Seed) truncateTables() {
	s.executeSqlFromFile(fmt.Sprintf("%stest_helpers/seed_scripts/truncate.sql", s.path))
}

func Connect() (*gorm.DB, error) {
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

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	conn, _ := db.DB()

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(20)
	conn.SetConnMaxIdleTime(time.Hour)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return db, nil
}
