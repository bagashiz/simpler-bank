package v1

import (
	"net/http"

	helper "github.com/bagashiz/simpler-bank/controllers"
	v1s "github.com/bagashiz/simpler-bank/services/v1"
	"github.com/gin-gonic/gin"
)

// CreateUser is a function for creating a new user.
func CreateUser(ctx *gin.Context) {
	var userService v1s.UserService

	if err := ctx.ShouldBindJSON(&userService.User); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}

	user, err := userService.CreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
