package v1services

import (
	"context"

	"github.com/bagashiz/simpler-bank/db"
	"github.com/bagashiz/simpler-bank/models"
)

// CreateTransferParams is a struct to hold the parameters for creating a transfer.
type CreateTransferParams struct {
	FromAccountID uint  `json:"from_account_id"`
	ToAccountID   uint  `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// CreateTransfer is a function to create a transfer between two accounts.
func CreateTransfer(ctx context.Context, arg CreateTransferParams) (models.Transfer, error) {
	transfer := models.Transfer{
		FromAccountID: arg.FromAccountID,
		ToAccountID:   arg.ToAccountID,
		Amount:        arg.Amount,
	}

	err := db.GetDB().Create(&transfer).Error

	return transfer, err
}

// GetTransfer is a function to get a transfer by transfer ID.
func GetTransfer(ctx context.Context, id uint) (models.Transfer, error) {
	var transfer models.Transfer
	err := db.GetDB().First(&transfer, id).Error

	return transfer, err
}

// ListTransfersParams is a struct to hold the parameters for listing transfers of a pair of accounts with pagination.
type ListTransfersParams struct {
	FromAccountID uint `json:"from_account_id"`
	ToAccountID   uint `json:"to_account_id"`
	Limit         int  `json:"limit"`
	Offset        int  `json:"offset"`
}

// ListTransfers is a function to list transfers of a pair of accounts with pagination.
func ListTransfers(ctx context.Context, arg ListTransfersParams) ([]models.Transfer, error) {
	var transfers []models.Transfer
	err := db.GetDB().
		Limit(arg.Limit).
		Offset(arg.Offset).
		Where("from_account_id = ? AND to_account_id = ?", arg.FromAccountID, arg.ToAccountID).
		Find(&transfers).
		Error

	return transfers, err
}
