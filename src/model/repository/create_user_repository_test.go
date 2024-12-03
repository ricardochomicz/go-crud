package repository

import (
	"os"
	"testing"

	"github.com/ricardochomicz/go-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	database_name := "users_database_test"
	collection_name := "users_collection_test"

	os.Setenv("MONGODB_DATABASE", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.ClearCollections()

	mtestDb.Run("when_sending_a_valid_user", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@email.com", "test", "test", 46,
		)
		userDomain, err := repo.CreateUser(domain)

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "test@email.com")
		assert.EqualValues(t, userDomain.GetName(), "test")
		assert.EqualValues(t, userDomain.GetAge(), int8(46))
		assert.EqualValues(t, userDomain.GetPassword(), "test")
	})

	mtestDb.Run("return_error_when_sending_an_invalid_user", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@email.com", "test", "test", 46,
		)
		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})
}
