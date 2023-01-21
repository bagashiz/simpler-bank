package models

// Transfers entity struct.
type Transfer struct {
	Model
	FromAccountID uint  `gorm:"not null;index;index:from_to_accounts_idx" json:"from_account_id"`
	ToAccountID   uint  `gorm:"not null;index;index:from_to_accounts_idx" json:"to_account_id"`
	Amount        int64 `gorm:"not null;comment:must be positive" json:"amount"`
}
