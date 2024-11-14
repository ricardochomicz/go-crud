package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardochomicz/go-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/userCreate", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
