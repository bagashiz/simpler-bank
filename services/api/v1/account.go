package v1services

import (
	"context"

	"github.com/bagashiz/simpler-bank/db"
	"github.com/bagashiz/simpler-bank/models"
)

// CreateAccountParams is a struct to hold the parameters for creating an account.
type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

// CreateAccount is a function to create an account for a user.
func CreateAccount(ctx context.Context, arg CreateAccountParams) (models.Account, error) {
	account := models.Account{
		Owner:    arg.Owner,
		Balance:  arg.Balance,
		Currency: arg.Currency,
	}
	err := db.GetDB().Create(&account).Error

	return account, err
}

// GetAccount is a function to get an account for a user by account ID.
func GetAccount(ctx context.Context, id uint) (models.Account, error) {
	var account models.Account
	err := db.GetDB().First(&account, id).Error

	return account, err
}

// getAccountForUpdate is a function to get an account for a user by account ID for update using FOR NO KEY UPDATE feature.
func getAccountForUpdate(ctx context.Context, id uint) (models.Account, error) {
	var account models.Account
	err := db.GetDB().
		Set("gorm:query_option", "FOR NO KEY UPDATE").
		First(&account, id).
		Error

	return account, err
}

// ListAccountsParams is a struct to hold the parameters for listing accounts for a user with pagination.
type ListAccountsParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// ListAccounts is a function to list accounts for a user with pagination.
func ListAccounts(ctx context.Context, arg ListAccountsParams) ([]models.Account, error) {
	var accounts []models.Account
	err := db.GetDB().Limit(arg.Limit).Offset(arg.Offset).Find(&accounts).Error

	return accounts, err
}

// DeleteAccount is a function to delete an account for a user by account ID.
func DeleteAccount(ctx context.Context, id uint) error {
	err := db.GetDB().Delete(&models.Account{}, id).Error

	return err
}

// UpdateAccountBalanceParams is a struct to hold the parameters for changing the balance of an account.
type UpdateAccountBalanceParams struct {
	ID     uint  `json:"id"`
	Amount int64 `json:"amount"`
}

// UpdateAccountBalance is a function to change the balance of an account.
func UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (models.Account, error) {
	var account models.Account
	var updatedBalance int64

	account, _ = getAccountForUpdate(ctx, arg.ID)
	updatedBalance = account.Balance + arg.Amount

	err := db.GetDB().Model(&account).Update("balance", updatedBalance).Error

	return account, err
}
