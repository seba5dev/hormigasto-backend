package models

import "time"

type Transaction struct {
	ID           uint                 `gorm:"primaryKey" json:"id"`
	Amount       float64              `gorm:"type:decimal(10,2);not null" json:"amount"`
	Date         time.Time            `gorm:"not null" json:"date"`
	AccountInID  *uint                `json:"account_in_id"`
	AccountOutID *uint                `json:"account_out_id"`
	TypeID       uint                 `gorm:"not null" json:"type_id"`
	CategoryID   *uint                `json:"category_id"`
	Description  string               `json:"description"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	AccountIn    *Account             `gorm:"foreignKey:AccountInID" json:"-"`
	AccountOut   *Account             `gorm:"foreignKey:AccountOutID" json:"-"`
	Type         TransactionType      `gorm:"foreignKey:TypeID" json:"-"`
	Category     *TransactionCategory `gorm:"foreignKey:CategoryID" json:"-"`
}
