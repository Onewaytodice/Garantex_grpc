package garantex

import (
	"Garantex_grpc/config"
	"Garantex_grpc/internal/domain"
	"context"
	"go.uber.org/zap"
	"io"
	"net/http"
)

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
	return depth.ToDomain()
}
