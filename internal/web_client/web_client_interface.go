package webclient

import (
	"Garantex_grpc/internal/domain"
	"context"
)

//go:generate mockgen -source=web_client_interface.go -destination=./mocks/mock_web_client_interface.go -package=mocks
type WebClient interface {
	GetRatesFromDepth(ctx context.Context, market string) (domain.Rates, error)
}
