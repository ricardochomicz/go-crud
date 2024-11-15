package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardochomicz/go-crud/src/model/service"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

// FindUserByID implements UserControllerInterface.
func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	panic("unimplemented")
}
