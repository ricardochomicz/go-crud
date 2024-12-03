package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/ricardochomicz/go-crud/src/configuration/logger"
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	logger.Info("Init findAllUser controller", zap.String("journey", "findAllUser"))

	// Chama o serviço para buscar todos os usuários
	userDomain, err := uc.service.FindAllUsersService()
	if err != nil {
		logger.Error("Error trying to find all users service",
			err,
			zap.String("journey", "findAllUser Service"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindAllUser  controller executed successfully",
		zap.String("journey", "findAllUser "))

	// Converte os usuários do domínio para a resposta e retorna
	c.JSON(http.StatusOK, view.ConvertDomainsToResponses(userDomain))
}

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserById controller", zap.String("journey", "findUserById"))

	userId := c.Param("userId")

	logger.Info("User authenticated successfully")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to parse user id",
			err,
			zap.String("journey", "findUserByIdService"))
		errorMessage := rest_err.NewBadRequestError("invalid user id")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Error("Error trying to find user by id service",
			err,
			zap.String("journey", "findUserByIdService"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed successfully",
		zap.String("journey", "findUserById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "findUserByEmail"))

	email := c.Param("userEmail")

	logger.Info(fmt.Sprintf("User authenticated successfully"))

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error trying to parse user id",
			err,
			zap.String("journey", "findUserByEmailService"))
		errorMessage := rest_err.NewBadRequestError("invalid user email")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(email)
	if err != nil {
		logger.Error("Error trying to find user by email service",
			err,
			zap.String("journey", "findUserByEmailService"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
