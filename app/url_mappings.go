package app

import (
	"github.com/MicahJackson/go-url-shortener/handler"
	"github.com/gin-gonic/gin"
)

func MapRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Wussup homie.",
		})
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.RedirectShortUrl(c)
	})
}
