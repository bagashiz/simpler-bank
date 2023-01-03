package models

import "time"

type Model struct {
	ID        uint      `gorm:"not null;primaryKey" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
