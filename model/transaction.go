package model

import "time"

type Transaction struct {
	ID           string    `json:"id" xorm:"'id' pk"`
	AccountID    string    `json:"account_id" xorm:"'account_id'"`
	Amount       float64   `json:"amount" xorm:"'amount'"`
	MerchantName string    `json:"merchant_name" xorm:"'merchant_name'"`
	Description  string    `json:"description" xorm:"'description'"`
	Status       string    `json:"status" xorm:"'status'"`
	CreatedAt    time.Time `json:"created_at" xorm:"'created_at' created"`
}
