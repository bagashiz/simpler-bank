package models

// Entries entity struct.
type Entry struct {
	Model
	AccountID uint  `gorm:"not null;index;foreignKey" json:"account_id"`
	Amount    int64 `gorm:"not null;comment:can be negative or positive" json:"amount"`
}
