package storage

import (
	"Garantex_grpc/internal/domain"
	"context"
)

//go:generate mockgen -source=storage_interface.go -destination=./mocks/mock_storage_interface.go -package=mocks
type Storager interface {
	SaveRates(ctx context.Context, rates domain.Rates) error
}
