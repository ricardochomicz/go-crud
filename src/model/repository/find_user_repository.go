package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/ricardochomicz/go-crud/src/configuration/logger"
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
	"github.com/ricardochomicz/go-crud/src/model/repository/entity"
	"github.com/ricardochomicz/go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findAllUsers repository", zap.String("journey", "findAllUsers"))

	collection_name := os.Getenv(MONGODB_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		errorMessage := "Error trying to find all users"
		logger.Error(errorMessage, err, zap.String("journey", "findAllUsers"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	var userEntities []entity.UserEntity
	if err = cursor.All(context.Background(), &userEntities); err != nil {
		errorMessage := "Error trying to convert to user entities"
		logger.Error(errorMessage, err, zap.String("journey", "findAllUsers"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindAllUsers repository executed successfully",
		zap.String("journey", "findAllUsers"))

	return converter.ConverterEntitiesToDomain(userEntities), nil
}

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByEmail repository", zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(
		context.Background(),
		filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User with email %s not found", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConverterEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserById repository", zap.String("journey", "findUserById"))

	collection_name := os.Getenv(MONGODB_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(
		context.Background(),
		filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User with id %s not found", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserById"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("findUserById repository executed successfully",
		zap.String("journey", "findUserById"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConverterEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"))

	collection_name := os.Getenv(MONGODB_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}

	err := collection.FindOne(
		context.Background(),
		filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User with email or password not found"
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
			return nil, rest_err.NewForbidenError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConverterEntityToDomain(*userEntity), nil
}
