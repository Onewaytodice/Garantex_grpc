package postgres

import (
	"Garantex_grpc/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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
	q := `INSERT INTO rates (timestamp, ask, bid) VALUES ($1, $2, $3)`

	_, err := s.db.ExecContext(ctx, q, rates.Timestamp, rates.AskPrice, rates.BidPrice)
	if err != nil {
		s.logger.Debug("postgres insert rates error", zap.Error(err))
		return err
	}

	return nil
}
