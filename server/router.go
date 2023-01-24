package server

import (
	"log"
	"net/http"

	v1c "github.com/bagashiz/simpler-bank/controllers/api/v1"
	"github.com/bagashiz/simpler-bank/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var router *gin.Engine

// SetupRouter performs all route operations.
func SetupRouter() {
	router = gin.Default()

	// register custom validator
	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("currency", helpers.ValidCurrency)
	}

	v1 := router.Group("/api/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Simpler Bank API V1",
		})
	})

	// accounts routes
	v1.POST("/accounts", v1c.CreateAccount)
	v1.GET("/accounts/:id", v1c.GetAccount)
	v1.GET("/accounts", v1c.ListAccounts)
}

// Start attaches the router to a server and starts listening and serving HTTP requests from specified address.
func Start(address string) {
	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
