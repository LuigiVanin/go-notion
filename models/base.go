package models

import "time"

type Base struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
