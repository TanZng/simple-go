package pet

import (
	"context"
	"errors"
	"testing"

	"simple-go/validation"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PetServiceTestSuite struct {
	suite.Suite
	mockCtrl     *gomock.Controller
	petStoreMock *MockPetStoreInterface
	service      PetServiceT
	name         string
	kind         string
}

func (suite *PetServiceTestSuite) SetupSuite() {
	suite.mockCtrl = gomock.NewController(suite.T())

}

func (suite *PetServiceTestSuite) SetupTest() {
	suite.petStoreMock = NewMockPetStoreInterface(suite.mockCtrl)
	suite.service = NewPetService(suite.petStoreMock)
	suite.name = "Toby"
	suite.kind = "Dog"
}

func (suite *PetServiceTestSuite) TearDownSuite() {
	suite.mockCtrl.Finish()
}

func (suite *PetServiceTestSuite) TestGetByID() {

	suite.T().Run("When pet does exists", func(t *testing.T) {
		id := uuid.New()

		suite.petStoreMock.EXPECT().GetByID(id).Return(&Pet{
			ID:   id,
			Name: suite.name,
			Kind: suite.kind,
		}, nil)

		petService := NewPetService(suite.petStoreMock)

		pet, err := petService.GetByID(context.Background(), id)

		assert.NoError(t, err)
		assert.Equal(t, pet.ID, id)
		assert.Equal(t, pet.Name, suite.name)
		assert.Equal(t, pet.Kind, suite.kind)

	})

	suite.T().Run("When pet does not exists", func(t *testing.T) {
		id := uuid.New()

		errorMessage := errors.New(validation.RESOURCE_NOT_FOUND_WITH_ID)

		suite.petStoreMock.EXPECT().GetByID(id).Return(nil, errorMessage)

		petService := NewPetService(suite.petStoreMock)

		pet, err := petService.GetByID(context.Background(), id)

		assert.Error(t, err)
		assert.Equal(t, err, errorMessage)
		assert.Nil(t, pet)
	})

}

func (suite *PetServiceTestSuite) TestCreate() {
	newUUID := uuid.New()

	suite.T().Run("When create pet succeeds", func(t *testing.T) {
		pet := &Pet{
			Name: suite.name,
			Kind: suite.kind,
		}

		suite.petStoreMock.EXPECT().Create(pet).Return(newUUID, nil)

		petService := NewPetService(suite.petStoreMock)

		createdID, err := petService.Create(context.Background(), pet)

		assert.NoError(t, err)
		assert.Equal(t, pet.ID, createdID)
	})

	suite.T().Run("When create pet fails", func(t *testing.T) {
		invalidID := uuid.UUID{}
		pet := &Pet{
			Name: suite.name,
			Kind: suite.kind,
		}
		errorMessage := errors.New(validation.ERROR_MESSAGE)

		suite.petStoreMock.EXPECT().Create(pet).Return(invalidID, errorMessage)

		petService := NewPetService(suite.petStoreMock)

		createdID, err := petService.Create(context.Background(), pet)

		assert.Error(t, err)
		assert.Equal(t, createdID, invalidID)
		assert.Equal(t, errorMessage, err)
	})

}

func TestPetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PetServiceTestSuite))
}
