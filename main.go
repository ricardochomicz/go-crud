package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ricardochomicz/go-crud/src/configuration/database/mongodb"
	"github.com/ricardochomicz/go-crud/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal(
			"Error trying to connect to database",
			err.Error())
		return
	}

	userController := initDependecies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8086"); err != nil {
		log.Fatal(err)
	}

}
