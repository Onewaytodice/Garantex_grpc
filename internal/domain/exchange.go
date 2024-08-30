package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type Rates struct {
	ID        int
	Market    string
	Timestamp time.Time
	AskPrice  decimal.Decimal
	BidPrice  decimal.Decimal
}
