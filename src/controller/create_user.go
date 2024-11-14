package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ricardochomicz/go-crud/src/configuration/rest_err"
	"github.com/ricardochomicz/go-crud/src/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("invalid json body: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

}
