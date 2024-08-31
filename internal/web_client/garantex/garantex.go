package garantex

import (
	"Garantex_grpc/internal/config"
	"Garantex_grpc/internal/domain"
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	getRatesFromDepthDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Subsystem: "garantex_web_client",
		Name:      "get_rates_from_depth_duration_seconds",
		Help:      "Request to external API duration in seconds",
		Buckets:   prometheus.DefBuckets,
	})
	getRatesFromDepthRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "",
		Subsystem:   "garantex_web_client",
		Name:        "get_rates_from_depth_requests_total",
		Help:        "Total requests to external API partition by status",
		ConstLabels: nil,
	}, []string{"status"})

	tracer = otel.Tracer("garantex")
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
	// Tracing
	ctx, span := tracer.Start(ctx, "GetRatesFromDepth")
	defer span.End()
	// Metrics
	start := time.Now()
	defer func() {
		getRatesFromDepthDuration.Observe(time.Since(start).Seconds())
	}()

	market = strings.ToLower(market)
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

	getRatesFromDepthRequests.WithLabelValues(resp.Status).Inc()

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
