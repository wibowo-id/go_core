package entity

import (
	"math/big"
	"time"
)

type Account struct {
	Id        big.Int   `gorm:"primaryKey:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Code      string    `gorm:"type:varchar(255)" json:"code"`
	IsDebt    bool      `gorm:"type:bool" json:"is_debt"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}
