package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ricardochomicz/go-crud/src/controller"
	"github.com/ricardochomicz/go-crud/src/controller/routes"
	"github.com/ricardochomicz/go-crud/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//initial dependences
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8086"); err != nil {
		log.Fatal(err)
	}

}
