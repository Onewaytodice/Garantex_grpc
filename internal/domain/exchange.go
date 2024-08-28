package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type Rates struct {
	ID        int
	Timestamp time.Time
	AskPrice  decimal.Decimal
	BidPrice  decimal.Decimal
}
