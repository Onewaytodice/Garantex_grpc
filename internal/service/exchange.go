package service

import (
	"Garantex_grpc/internal/domain"
	"Garantex_grpc/internal/storage"
	webclient "Garantex_grpc/internal/web_client"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

//go:generate mockgen -source=exchange.go -destination=./mocks/mock_exchange_interface.go -package=mocks
type Exchanger interface {
	GetAndSaveRates(ctx context.Context, market string) (domain.Rates, error)
}

var ErrGetRates = errors.New("failed to get rates")

type Exchange struct {
	logger  *zap.Logger
	web     webclient.WebClient
	storage storage.Storager
}

func NewExchange(logger *zap.Logger, web webclient.WebClient, storage storage.Storager) *Exchange {
	return &Exchange{
		logger:  logger,
		web:     web,
		storage: storage,
	}
}

func (e *Exchange) GetAndSaveRates(ctx context.Context, market string) (domain.Rates, error) {
	rates, err := e.web.GetRatesFromDepth(ctx, market)
	if err != nil {
		e.logger.Error("failed to get rates from source", zap.Error(err))
		return domain.Rates{}, ErrGetRates
	}

	e.logger.Debug(fmt.Sprintf("got rates: %v", rates))

	// strong consistency between database data and responses
	err = e.storage.SaveRates(ctx, rates)
	if err != nil {
		e.logger.Error("failed to save rates", zap.Error(err))
		return domain.Rates{}, ErrGetRates
	}
	return rates, nil
}
