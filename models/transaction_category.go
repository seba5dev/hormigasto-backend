package models

import "time"

type TransactionCategory struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `gorm:"not null" json:"name"`
	TypeID      uint            `gorm:"not null" json:"type_id"`
	Color       string          `json:"color"`
	Icon        string          `json:"icon"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	Type        TransactionType `gorm:"foreignKey:TypeID" json:"-"`
}
