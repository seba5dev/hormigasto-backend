package models

import "time"

type TransactionType struct {
	ID          uint                  `gorm:"primaryKey" json:"id"`
	Name        string                `gorm:"not null" json:"name"`
	Color       string                `json:"color"`
	Icon        string                `json:"icon"`
	Description string                `json:"description"`
	CreatedAt   time.Time             `json:"created_at"`
	Categories  []TransactionCategory `gorm:"foreignKey:TypeID" json:"categories,omitempty"`
}
