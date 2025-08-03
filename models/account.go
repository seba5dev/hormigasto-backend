package models

import "time"

type Account struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	UserID       uint        `gorm:"not null" json:"user_id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	TypeID       uint        `gorm:"not null" json:"type_id"`
	CurrencyCode string      `gorm:"not null" json:"currency_code"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	User         User        `gorm:"foreignKey:UserID" json:"-"`
	Type         AccountType `gorm:"foreignKey:TypeID" json:"-"`
}
