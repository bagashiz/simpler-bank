package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// SetupRouter performs all route operations.
func SetupRouter() {
	router = gin.Default()
	v1 := router.Group("/api/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Simpler Bank API V1",
		})
	})
}

// Start attaches the router to a server and starts listening and serving HTTP requests from specified address.
func Start(address string) {
	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
