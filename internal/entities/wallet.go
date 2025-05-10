package entities

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	BaseGorm
	UserID     uuid.UUID       `json:"-"`
	CurrencyID uuid.UUID       `json:"currency"`
	Balance    decimal.Decimal `json:"balance" gorm:"type:numeric(20,6)"`
}
