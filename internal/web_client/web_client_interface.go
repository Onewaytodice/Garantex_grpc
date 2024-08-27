package webclient

import "Garantex_grpc/internal/domain"

type WebClient interface {
	GetRatesFromDepth(market string) domain.Rates
}
