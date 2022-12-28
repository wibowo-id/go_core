package entity

import "math/big"

type Product struct {
	Id          big.Int `gorm:"primaryKey:auto_increment" json:"id"`
	Name        string  `gorm:"type:varchar(255)" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	AccountId   big.Int `json:"account_id"`
}
