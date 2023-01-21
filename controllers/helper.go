package helper

import "github.com/gin-gonic/gin"

// ErrorResponse is a common format for API errors.
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
