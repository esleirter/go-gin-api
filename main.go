package main

import (
	"go-gin-api/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/dummy", func(c *gin.Context) {
		c.String(http.StatusOK, "Hellow everyone 23!")
	})

	router.GET("/bye", func(c *gin.Context) {
		c.String(http.StatusOK, "Bye world, good luck")
	})

	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + configs.PORT)

}
