package server

import (
	"github.com/bagashiz/simpler-bank/configs"
	v1c "github.com/bagashiz/simpler-bank/controllers/api/v1"
	"github.com/bagashiz/simpler-bank/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config configs.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server.
func NewServer(config configs.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	// register custom validator
	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("currency", helpers.ValidCurrency)
	}

	server.setupRouter()

	return server, nil
}

// setupRouter performs all route operations.
func (server *Server) setupRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	// v1 accounts routes
	v1.POST("/accounts", v1c.CreateAccount)
	v1.GET("/accounts/:id", v1c.GetAccount)
	v1.GET("/accounts", v1c.ListAccounts)

	// v1 transfers routes
	v1.POST("/transfers", v1c.CreateTransfer)

	server.router = router
}

// Start attaches the router to a server and starts listening and serving HTTP requests from specified address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
