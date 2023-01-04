package models

// Accounts entity struct.
type Account struct {
	Model
	Owner         string     `gorm:"type:varchar;not null;index;uniqueIndex:owner_currency_key" json:"owner"`
	Balance       int64      `gorm:"not null" json:"balance"`
	Currency      string     `gorm:"type:varchar(3);not null;uniqueIndex:owner_currency_key" json:"currency"`
	TransfersFrom []Transfer `gorm:"foreignKey:FromAccountID;references:ID"`
	TransfersTo   []Transfer `gorm:"foreignKey:ToAccountID;references:ID"`
	Entries       []Entry    `gorm:"foreignKey:AccountID;references:ID"`
}
