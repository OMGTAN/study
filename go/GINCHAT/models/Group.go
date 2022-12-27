package models

import "time"

// swagger:model Message
type Group struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	name      string
}
