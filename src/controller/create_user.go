package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardochomicz/go-crud/src/configuration/logger"
	"github.com/ricardochomicz/go-crud/src/configuration/validation"
	"github.com/ricardochomicz/go-crud/src/model/request"
	"github.com/ricardochomicz/go-crud/src/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init create user controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "Teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("Create user successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)

}
