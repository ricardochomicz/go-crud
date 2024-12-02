package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/ricardochomicz/go-crud/src/configuration/logger"
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model"
	"github.com/ricardochomicz/go-crud/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// func (uc *userControllerInterface) FindUser(c *gin.Context) {
// }

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserById controller", zap.String("journey", "findUserById"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated successfully: %#v", user))

	userId := c.Param("userId")

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

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated successfully: %#v", user))

	email := c.Param("userEmail")

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
