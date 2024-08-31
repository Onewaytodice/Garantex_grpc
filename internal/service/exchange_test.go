package service

import (
	"Garantex_grpc/internal/domain"
	storagemock "Garantex_grpc/internal/storage/mocks"
	webmock "Garantex_grpc/internal/web_client/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"testing"
	"time"
)

type webmockBehavior func(m *webmock.MockWebClient)
type storagemockBehaivor func(m *storagemock.MockStorager)

var (
	logger = zap.NewNop()
	ctx    = context.Background()
)

func TestExchange_GetAndSaveRates(t *testing.T) {
	market := "usdt"
	rates := domain.Rates{
		Timestamp: time.Unix(100000, 0),
		AskPrice:  decimal.NewFromFloat(100.200),
		BidPrice:  decimal.NewFromFloat(300.400),
	}
	tests := []struct {
		name        string
		webMock     webmockBehavior
		storageMock storagemockBehaivor
		want        domain.Rates
		wantErr     bool
	}{
		{
			name: "test ok",
			webMock: func(m *webmock.MockWebClient) {
				m.EXPECT().GetRatesFromDepth(ctx, market).Return(rates, nil)
			},
			storageMock: func(m *storagemock.MockStorager) {
				m.EXPECT().SaveRates(ctx, rates).Return(nil)
			},
			want:    rates,
			wantErr: false,
		},
		{
			name: "test get rates error",
			webMock: func(m *webmock.MockWebClient) {
				m.EXPECT().GetRatesFromDepth(ctx, market).Return(domain.Rates{}, errors.New("get rates from web error"))
			},
			storageMock: func(m *storagemock.MockStorager) {},
			want:        domain.Rates{},
			wantErr:     true,
		},
		{
			name: "test save rates error",
			webMock: func(m *webmock.MockWebClient) {
				m.EXPECT().GetRatesFromDepth(ctx, market).Return(rates, nil)
			},
			storageMock: func(m *storagemock.MockStorager) {
				m.EXPECT().SaveRates(ctx, rates).Return(errors.New("save rates to db error"))
			},
			want:    domain.Rates{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()
			mockWebCli := webmock.NewMockWebClient(c)
			mockStorage := storagemock.NewMockStorager(c)

			e := NewExchange(logger, mockWebCli, mockStorage)
			tt.webMock(mockWebCli)
			tt.storageMock(mockStorage)

			got, err := e.GetAndSaveRates(ctx, market)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAndSaveRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAndSaveRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}
