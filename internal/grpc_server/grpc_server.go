package grpcserver

import (
	"Garantex_grpc/internal/service"
	pbexchange "Garantex_grpc/proto/exchange_v1"
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"time"
)

var (
	getRatesRequests = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: "gRPC_server",
		Name:      "get_rates_requests_total",
		Help:      "Total number of request",
	})
	getRatesDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Subsystem: "gRPC_server",
		Name:      "get_rates_duration_seconds",
		Help:      "Duration of request",
		Buckets:   prometheus.DefBuckets,
	})

	tracer = otel.Tracer("grpc_server")
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
	// Tracing
	ctx, span := tracer.Start(ctx, "GetRates")
	defer span.End()
	// Metrics
	getRatesRequests.Inc()
	start := time.Now()
	defer func() {
		getRatesDuration.Observe(time.Since(start).Seconds())
	}()

	rates, err := e.exchange.GetAndSaveRates(ctx, req.GetMarket().String())
	if err != nil {
		return &pbexchange.GetRatesResponse{}, err
	}
	return &pbexchange.GetRatesResponse{
		Timestamp: rates.Timestamp.Unix(),
		Market:    req.GetMarket(),
		Ask:       rates.AskPrice.InexactFloat64(),
		Bid:       rates.BidPrice.InexactFloat64(),
	}, nil
}
