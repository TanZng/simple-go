package pet

import (
	"errors"
	"regexp"
	"testing"

	"simple-go/permissions/test_helpers"
	"simple-go/validation"

	"github.com/DATA-DOG/go-sqlmock"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PetRepositoryTestSuite struct {
	suite.Suite

	mockCtrl *gomock.Controller
	db       test_helpers.MockedGORM
	columns  []string

	pet   Pet
	name  string
	kind  string
	petId uuid.UUID

	petSelectQuery string
	petInsertQuery string

	repository PetRepositoryT
}

func (suite *PetRepositoryTestSuite) SetupSuite() {
	suite.mockCtrl = gomock.NewController(suite.T())
	petId := uuid.New()
	name := "Jane"
	kind := "Dog"

	suite.petSelectQuery = "SELECT * FROM `pets` WHERE `pets`.`id` = ? AND `pets`.`deleted_at` IS NULL"
	suite.petInsertQuery = "INSERT INTO `pets` (`name`,`kind`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?)"

	suite.columns = []string{"id", "name", "kind"}

	suite.pet = Pet{
		ID:   petId,
		Name: name,
		Kind: kind,
	}
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite *PetRepositoryTestSuite) SetupTest() {
	suite.db = test_helpers.MockGORM()
	suite.repository, _ = NewPetRepository(suite.db)
}

func (suite *PetRepositoryTestSuite) TearDownSuite() {
	suite.mockCtrl.Finish()
}

func (suite *PetRepositoryTestSuite) TearDownTest() {
	suite.db.Close()
}

func (suite *PetRepositoryTestSuite) TestGetByID() {

	suite.T().Run("When pet does exists", func(t *testing.T) {
		expectedPet := &suite.pet

		// Mock Rows and Register (Row)
		petRow := sqlmock.NewRows(suite.columns).
			AddRow(
				suite.pet.ID,
				suite.pet.Name,
				suite.pet.Kind,
			)

		suite.db.Mock().
			ExpectQuery(regexp.
				QuoteMeta(suite.petSelectQuery)).
			WithArgs(suite.pet.ID).
			WillReturnRows(petRow)

		actualPet, err := suite.repository.GetByID(suite.pet.ID)

		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), expectedPet, actualPet)
	})

	suite.T().Run("When pet does not exists", func(t *testing.T) {

		// Mock Rows and NO Register (Row)
		petRow := sqlmock.NewRows(suite.columns)

		suite.db.Mock().
			ExpectQuery(regexp.
				QuoteMeta(suite.petSelectQuery)).
			WithArgs(suite.pet.ID).
			WillReturnRows(petRow)

		pet, err := suite.repository.GetByID(suite.pet.ID)

		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), pet)
	})

}

func (suite *PetRepositoryTestSuite) TestCreate() {

	suite.T().Run("When create pet succeeds", func(t *testing.T) {
		pet := Pet{
			Name: suite.name,
			Kind: suite.kind,
		}

		suite.db.Mock().ExpectBegin()
		suite.db.Mock().
			ExpectExec(regexp.QuoteMeta(suite.petInsertQuery)).
			WithArgs(
				pet.Name,
				pet.Kind,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(0, 0))
		suite.db.Mock().ExpectCommit()

		createdID, err := suite.repository.Create(&pet)

		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), pet.ID, createdID)
	})

	suite.T().Run("When create pet fails", func(t *testing.T) {
		emptyID := uuid.UUID{}
		pet := Pet{
			Name: suite.pet.Name,
			Kind: suite.pet.Kind,
		}
		errorMessage := errors.New(validation.ERROR_MESSAGE)

		suite.db.Mock().ExpectBegin()
		suite.db.Mock().
			ExpectExec(regexp.QuoteMeta(suite.petInsertQuery)).
			WithArgs(
				pet.Name,
				pet.Kind,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnError(errorMessage)
		suite.db.Mock().ExpectCommit()

		createdID, err := suite.repository.Create(&pet)

		assert.Error(suite.T(), err)
		assert.Equal(suite.T(), emptyID, createdID)
		assert.Equal(suite.T(), emptyID, createdID)
	})

}

func TestPetRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PetRepositoryTestSuite))
}
