package webclient

import (
	"Garantex_grpc/internal/domain"
	"context"
)

type WebClient interface {
	GetRatesFromDepth(ctx context.Context, market string) (domain.Rates, error)
}
