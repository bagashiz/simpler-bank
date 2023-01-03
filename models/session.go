package models

import "time"

type Session struct {
	Model
	Username     string    `gorm:"type:varchar;not null" json:"username"`
	RefreshToken string    `gorm:"type:varchar;not null" json:"refresh_token"`
	UserAgent    string    `gorm:"type:varchar;not null" json:"user_agent"`
	ClientIP     string    `gorm:"type:varchar;not null" json:"client_ip"`
	IsBlocked    bool      `gorm:"not null;default:false" json:"is_blocked"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
}
