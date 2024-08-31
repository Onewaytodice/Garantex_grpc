package postgres

import (
	"Garantex_grpc/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"time"
)

var (
	saveRatesDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Subsystem: "postgreSQL",
		Name:      "save_rates_to_db_duration_seconds",
		Help:      "Database query duration in seconds",
		Buckets:   prometheus.ExponentialBuckets(0.001, 2, 10),
	})

	tracer = otel.Tracer("postgres")
)

type Storage struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewStorage(logger *zap.Logger, db *sqlx.DB) *Storage {
	return &Storage{
		logger: logger,
		db:     db,
	}
}

func (s *Storage) SaveRates(ctx context.Context, rates domain.Rates) error {
	// Tracing
	ctx, span := tracer.Start(ctx, "SaveRates")
	defer span.End()
	// Metrics
	start := time.Now()
	defer func() {
		saveRatesDuration.Observe(time.Since(start).Seconds())
	}()

	q := `INSERT INTO rates (timestamp, market, ask, bid) VALUES ($1, $2, $3, $4)`

	_, err := s.db.ExecContext(ctx, q, rates.Timestamp, rates.Market, rates.AskPrice, rates.BidPrice)
	if err != nil {
		s.logger.Debug("postgres insert rates error", zap.Error(err))
		return err
	}

	return nil
}
