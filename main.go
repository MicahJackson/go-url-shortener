package main

import (
	"fmt"
	"log"

	"github.com/MicahJackson/go-url-shortener/store"
	"github.com/MicahJackson/go-url-shortner/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvironmentVariables()

	router := gin.Default()
	mapRoutes(router)

	store.InitializeStore()

	serveAPI(router)
}

func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func mapRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.RedirectShortUrl(c)
	})
}

func serveAPI(router *gin.Engine) {
	err := router.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
