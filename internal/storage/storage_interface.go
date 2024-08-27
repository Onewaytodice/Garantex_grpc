package storage

import "Garantex_grpc/internal/domain"

type Storage interface {
	SaveRates(rates domain.Rates) error
}
