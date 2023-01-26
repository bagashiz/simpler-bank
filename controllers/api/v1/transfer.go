package v1controllers

import (
	"fmt"
	"net/http"

	"github.com/bagashiz/simpler-bank/helpers"
	v1s "github.com/bagashiz/simpler-bank/services/api/v1"
	"github.com/gin-gonic/gin"
)

// transferRequest body for creating a transfer transaction request.
type transferRequest struct {
	FromAccountID uint   `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   uint   `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

// CreateTransfer is a function to create a transfer transaction between two accounts.
func CreateTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	if valid := isValidAccountCurrency(ctx, req.FromAccountID, req.Currency); !valid {
		return
	}

	if valid := isValidAccountCurrency(ctx, req.ToAccountID, req.Currency); !valid {
		return
	}

	arg := v1s.TransferTxParams{
		CreateTransferParams: v1s.CreateTransferParams{
			FromAccountID: req.FromAccountID,
			ToAccountID:   req.ToAccountID,
			Amount:        req.Amount,
		},
	}

	result, err := v1s.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Response(result))
}

// isValidAccountCurrency is a function to check if the requested account and currency are the same as in the database.
func isValidAccountCurrency(ctx *gin.Context, accountID uint, currency string) bool {
	account, err := v1s.GetAccount(ctx, accountID)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return false
	}

	return true
}
