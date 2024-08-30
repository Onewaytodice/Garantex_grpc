package domain

import (
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	"time"
)

var ErrInvalidData = errors.New("invalid data")

type GrntxDepth struct {
	Timestamp int64       `json:"timestamp"`
	Asks      []GrntxRate `json:"asks"`
	Bids      []GrntxRate `json:"bids"`
}

type GrntxRate struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}

func UnmarshalGrntxDepth(data []byte) (GrntxDepth, error) {
	var r GrntxDepth
	err := json.Unmarshal(data, &r)
	return r, err
}

func (d *GrntxDepth) Valid() bool {
	switch {
	case d.Timestamp <= 0:
		return false
	case len(d.Asks) < 1:
		return false
	case len(d.Bids) < 1:
		return false
	default:
		return true
	}
}

func (d *GrntxDepth) ToDomain() (Rates, error) {
	if !d.Valid() {
		return Rates{}, ErrInvalidData
	}
	ask, err := decimal.NewFromString(d.Asks[0].Price)
	if err != nil {
		return Rates{}, ErrInvalidData
	}
	bid, err := decimal.NewFromString(d.Bids[0].Price)
	if err != nil {
		return Rates{}, ErrInvalidData
	}
	return Rates{
		Timestamp: time.Unix(d.Timestamp, 0),
		AskPrice:  ask,
		BidPrice:  bid,
	}, nil
}
