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
	ginRouter = gin.Default()
)

func RunApplication() {
	loadEnvironmentVariables()

	MapRoutes(ginRouter)

	store.InitializeStore()

	serveAPIOnPort(ginRouter, os.Getenv("LOCAL_PORT"))
}

func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveAPIOnPort(router *gin.Engine, port string) {
	err := router.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
