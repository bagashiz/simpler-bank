package v1controllers

import (
	"fmt"
	"net/http"

	"github.com/bagashiz/simpler-bank/helpers"
	v1s "github.com/bagashiz/simpler-bank/services/api/v1"
	"github.com/gin-gonic/gin"
)

// createAccountRequest is the request body for creating an account.
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

// CreateAccount is a function to create an account for a user.
func CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := v1s.CreateAccount(ctx, arg)
	if err != nil {
		if helpers.IsUniqueViolation(err) {
			ctx.JSON(http.StatusForbidden, helpers.ErrorResponse(fmt.Errorf("account with owner [%s] and currency [%s] already exists", req.Owner, req.Currency)))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Response(account))
}

// getAccountRequest is the request body for getting an account.
type getAccountRequest struct {
	ID uint `uri:"id" binding:"required"`
}

// GetAccount is a function to get an account for a user by account ID.
func GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	account, err := v1s.GetAccount(ctx, req.ID)
	if err != nil {
		// if record not found
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(fmt.Errorf("account with ID [%d] not found", req.ID)))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Response(account))
}

// listAccountRequest is the request body for listing accounts with pagination.
type listAccountRequest struct {
	PageID   int `form:"page_id" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=5,max=10"`
}

// ListAccounts is a function to list accounts for a user with pagination.
func ListAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := v1s.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Response(accounts))
}
