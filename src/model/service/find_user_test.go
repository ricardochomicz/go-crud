package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
	"github.com/ricardochomicz/go-crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_id_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"test@email.com", "test", "test", 46,
		)
		userDomain.SetID(id)

		repository.EXPECT().FindUserById(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIDService(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
	})

	t.Run("return_error_from_database", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().FindUserById(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByIDService(id)
		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_email_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@email.com"

		userDomain := model.NewUserDomain(
			email, "test", "test", 46,
		)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailService(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
	})

	t.Run("return_error_from_database", func(t *testing.T) {

		email := "test@email.com"
		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByEmailService(email)
		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_sending_a_valid_email_and_password_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@email.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain(
			email, password, "test", 46,
		)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findUserByEmailAndPasswordService(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
	})

	t.Run("return_error_from_database", func(t *testing.T) {

		email := "test@email.com"
		password := strconv.FormatInt(rand.Int63(), 10)
		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.findUserByEmailAndPasswordService(email, password)
		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}
