package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/ricardochomicz/go-crud/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "users_database_test"
	collection_name := "users_collection_test"

	err := os.Setenv("MONGODB_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.ClearCollections()

	mtestDb.Run("when_sending_a_valid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@email.com",
			Password: "test",
			Name:     "Test User",
			Age:      46,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), "test@email.com")
		assert.EqualValues(t, userDomain.GetName(), "Test User")
		assert.EqualValues(t, userDomain.GetAge(), int8(46))
		assert.EqualValues(t, userDomain.GetPassword(), "test")

	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, fmt.Sprintf("User with email %s not found", "test"))
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {
	database_name := "users_database_test"
	collection_name := "users_collection_test"

	err := os.Setenv("MONGODB_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.ClearCollections()

	mtestDb.Run("when_sending_a_valid_id_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@email.com",
			Password: "test",
			Name:     "Test User",
			Age:      46,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), "test@email.com")
		assert.EqualValues(t, userDomain.GetName(), "Test User")
		assert.EqualValues(t, userDomain.GetAge(), int8(46))
		assert.EqualValues(t, userDomain.GetPassword(), "test")

	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("test")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, fmt.Sprintf("User with id %s not found", "test"))
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	database_name := "users_database_test"
	collection_name := "users_collection_test"

	err := os.Setenv("MONGODB_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.ClearCollections()

	mtestDb.Run("when_sending_a_valid_email_and_password_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@email.com",
			Password: "test",
			Name:     "Test User",
			Age:      46,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), "test@email.com")
		assert.EqualValues(t, userDomain.GetName(), "Test User")
		assert.EqualValues(t, userDomain.GetAge(), int8(46))
		assert.EqualValues(t, userDomain.GetPassword(), "test")

	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "password")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "password")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
