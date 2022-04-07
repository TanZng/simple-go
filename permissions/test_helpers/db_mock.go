package test_helpers

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MockedGORM struct {
	sqlDB *sql.DB
	mock  sqlmock.Sqlmock
	gorm  *gorm.DB
}

func MockGORM() MockedGORM {
	sqlDB, mock, err := sqlmock.New()

	if err != nil {
		panic(err) // Error here
	}

	columns := []string{"version"}
	mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(
		mock.NewRows(columns).FromCSVString("1"),
	)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:       sqlDB,
		DriverName: "mysql",
	}), &gorm.Config{PrepareStmt: false})

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return MockedGORM{
		sqlDB: sqlDB,
		mock:  mock,
		gorm:  gormDB,
	}
}

func (m MockedGORM) Connect() (*gorm.DB, error) {
	return m.gorm, nil
}

func (m MockedGORM) Mock() sqlmock.Sqlmock {
	return m.mock
}

func (m MockedGORM) Close() {
	m.sqlDB.Close()
}
