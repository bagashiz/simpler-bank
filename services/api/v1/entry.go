package v1services

import (
	"context"

	"github.com/bagashiz/simpler-bank/db"
	"github.com/bagashiz/simpler-bank/models"
)

// CreateEntryParams is a struct to hold the parameters for creating an entry.
type CreateEntryParams struct {
	AccountID uint  `json:"account_id"`
	Amount    int64 `json:"amount"`
}

// CreateEntry is a function to create an entry for an account.
func CreateEntry(ctx context.Context, arg CreateEntryParams) (models.Entry, error) {
	entry := models.Entry{
		AccountID: arg.AccountID,
		Amount:    arg.Amount,
	}
	err := db.GetDB().Create(&entry).Error

	return entry, err
}

// GetEntry is a function to get an entry by entry ID.
func GetEntry(ctx context.Context, id uint) (models.Entry, error) {
	var entry models.Entry
	err := db.GetDB().First(&entry, id).Error

	return entry, err
}

// ListEntriesParams is a struct to hold the parameters for listing entries of an account with pagination.
type ListEntriesParams struct {
	AccountID uint `json:"account_id"`
	Limit     int  `json:"limit"`
	Offset    int  `json:"offset"`
}

// ListEntries is a function to list entries of an account with pagination.
func ListEntries(ctx context.Context, arg ListEntriesParams) ([]models.Entry, error) {
	var entries []models.Entry
	err := db.GetDB().
		Limit(arg.Limit).
		Offset(arg.Offset).
		Where("account_id = ?", arg.AccountID).
		Find(&entries).
		Error

	return entries, err
}
