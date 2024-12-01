package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct{
	TransactionID int64 `json:"transactionid"`
    Date time.Time `json:"date"`
    Amount decimal.Decimal `json:"amount"`
    Category string `json:"category"`
    Description string `json:"description"`
    PaymentMethod string `json:"payment_method"`
}