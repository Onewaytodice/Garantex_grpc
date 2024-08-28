package storage

import (
	"Garantex_grpc/internal/domain"
	"context"
)

type Storager interface {
	SaveRates(ctx context.Context, rates domain.Rates) error
}
