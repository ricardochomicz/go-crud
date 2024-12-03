package repository

import (
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPassword(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserById(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	DeleteUser(
		userId string,
	) *rest_err.RestErr
}
