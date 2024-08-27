package service

import (
	"Garantex_grpc/internal/domain"
	"Garantex_grpc/internal/storage"
	webclient "Garantex_grpc/internal/web_client"
)

type Exchanger interface {
	GetAndSaveRates() domain.Rates
}

type Exchange struct {
	web     webclient.WebClient
	storage storage.Storage
}
