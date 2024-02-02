package app

import (
	"fmt"
	"log"
	"os"

	"github.com/MicahJackson/go-url-shortener/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	Port = os.Getenv("LOCAL_PORT")
)

func StartApplication() {
	loadEnvironmentVariables()

	router := gin.Default()
	MapRoutes(router)

	store.InitializeStore()

	serveAPIOnPort(router, os.Getenv("LOCAL_PORT"))
}

func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveAPIOnPort(router *gin.Engine, Port string) {
	err := router.Run(":" + Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
