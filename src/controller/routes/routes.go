package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardochomicz/go-crud/src/controller"
	"github.com/ricardochomicz/go-crud/src/model"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/userCreate", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
}
