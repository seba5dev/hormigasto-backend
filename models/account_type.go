package models

import "time"

type AccountType struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Accounts    []Account `gorm:"foreignKey:TypeID" json:"accounts,omitempty"`
}
