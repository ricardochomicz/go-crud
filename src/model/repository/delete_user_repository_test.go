package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
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

	mtestDb.Run("when_sending_a_valid_userId_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)

		err := repo.DeleteUser("test")

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)

		err := repo.DeleteUser("test")

		assert.NotNil(t, err)

	})
}
