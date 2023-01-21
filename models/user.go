package models

import "time"

// User entity struct.
type User struct {
	Model
	Username          string    `gorm:"type:varchar;not null;unique;" json:"username"`
	FullName          string    `gorm:"type:varchar;not null" json:"full_name"`
	Email             string    `gorm:"type:varchar;not null;unique" json:"email"`
	HashedPassword    string    `gorm:"type:varchar;not null" json:"hashed_password"`
	PasswordChangedAt time.Time `gorm:"not null;autoCreateTime" json:"password_changed_at"`
	Accounts          []Account `gorm:"foreignKey:Owner;references:Username"`
	Sessions          []Session `gorm:"foreignKey:Username;references:Username"`
}
