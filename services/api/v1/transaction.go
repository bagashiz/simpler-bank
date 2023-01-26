package v1services

import (
	"context"

	"github.com/bagashiz/simpler-bank/db"
	"github.com/bagashiz/simpler-bank/models"
	"gorm.io/gorm"
)

// TransferTxParams contains the input parameters for the transfer transaction
// it contains the same parameters as CreateTransferParams
type TransferTxParams struct {
	CreateTransferParams
}

// TransferTxResult contains the result of the transfer transaction
type TransferTxResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount models.Account  `json:"from_account"`
	ToAccount   models.Account  `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

// TransferTx is a function that performs a money transfer from one account to another
// it creates a transfer record, add account entries, and update account balances in the database
func TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// create a transfer
		result.Transfer, err = CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// create an entry for the "from" account
		result.FromEntry, err = CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// add an entry for the "to" account
		result.ToEntry, err = CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// update account balances
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = updateAccountsBalance(
				ctx,
				arg.FromAccountID,
				-arg.Amount,
				arg.ToAccountID,
				arg.Amount,
			)
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = updateAccountsBalance(
				ctx,
				arg.ToAccountID,
				arg.Amount,
				arg.FromAccountID,
				-arg.Amount,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

// updateAccountsBalance is a function that updates the balance of two accounts in transaction.
// it is used to avoid deadlocks when updating account balances.
func updateAccountsBalance(
	ctx context.Context,
	accountID1 uint,
	amount1 int64,
	accountID2 uint,
	amount2 int64,
) (account1 models.Account, account2 models.Account, err error) {
	account1, err = UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})

	return
}
