package grpcserver

import (
	"Garantex_grpc/internal/service"
	pbexchange "Garantex_grpc/proto/exchange_v1"
	"context"
	"go.uber.org/zap"
)

type Exchange struct {
	logger   *zap.Logger
	exchange service.Exchanger

	pbexchange.UnimplementedExchangeGRPCServer
}

func NewExchange(logger *zap.Logger, exchange service.Exchanger) *Exchange {
	return &Exchange{
		logger:   logger,
		exchange: exchange,

		UnimplementedExchangeGRPCServer: pbexchange.UnimplementedExchangeGRPCServer{},
	}
}

func (e *Exchange) GetRates(ctx context.Context, req *pbexchange.GetRatesRequest) (*pbexchange.GetRatesResponse, error) {
	rates, err := e.exchange.GetAndSaveRates(ctx, req.GetMarket().String())
	if err != nil {
		return &pbexchange.GetRatesResponse{}, err
	}
	return &pbexchange.GetRatesResponse{
		Timestamp: rates.Timestamp.Unix(),
		Ask:       rates.AskPrice.InexactFloat64(),
		Bid:       rates.BidPrice.InexactFloat64(),
	}, nil
}
