package garantex

import (
	"Garantex_grpc/internal/config"
	"Garantex_grpc/internal/domain"
	"context"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

var tracer = otel.Tracer("garantex")

type Garantex struct {
	logger *zap.Logger
	client *http.Client
	url    string
}

func NewGarantex(logger *zap.Logger, config config.GarantexConfig) *Garantex {
	return &Garantex{
		logger: logger,
		client: &http.Client{
			Timeout: config.Timeout,
		},
		url: config.URL,
	}
}

func (g *Garantex) GetRatesFromDepth(ctx context.Context, market string) (domain.Rates, error) {
	market = strings.ToLower(market)

	ctx, span := tracer.Start(ctx, "GetRatesFromDepth")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, "GET", g.url+"/depth?market="+market, nil)
	if err != nil {
		g.logger.Debug("garantex new request error", zap.Error(err))
		return domain.Rates{}, err
	}
	resp, err := g.client.Do(req)
	if err != nil {
		g.logger.Debug("garantex do request error", zap.Error(err))
		return domain.Rates{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.logger.Debug("garantex read body error", zap.Error(err))
		return domain.Rates{}, err
	}
	depth, err := domain.UnmarshalGrntxDepth(body)
	if err != nil {
		g.logger.Debug("garantex unmarshal body error", zap.Error(err))
		return domain.Rates{}, err
	}
	rates, err := depth.ToDomain()
	if err != nil {
		g.logger.Debug("garantex to domain error", zap.Error(err))
		return domain.Rates{}, err
	}
	rates.Market = market
	return rates, nil
}
