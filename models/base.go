package models

import "time"

// Model includes fields ID and CreatedAt that can be embeded into another entity struct.
type Model struct {
	ID        uint      `gorm:"not null;primaryKey" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
