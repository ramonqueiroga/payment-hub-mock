package model

import (
	"github.com/jinzhu/gorm"
)

//TransactionModel define the database transaction that contains the payment information
type TransactionModel struct {
	gorm.Model
	PaymentID string  `json:"payment_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}
