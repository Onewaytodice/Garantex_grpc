package grpcserver

import (
	"Garantex_grpc/internal/domain"
	exchangemock "Garantex_grpc/internal/service/mocks"
	pbexchange "Garantex_grpc/proto/exchange_v1"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"reflect"
	"testing"
	"time"
)

type exchangeBehaivor func(m *exchangemock.MockExchanger)

var (
	logger = zap.NewNop()
	ctx    = context.Background()
)

func TestExchange_GetRates(t *testing.T) {
	market := "usdtrub"
	rates := domain.Rates{
		Timestamp: time.Unix(100000, 0),
		AskPrice:  decimal.NewFromFloat(100.200),
		BidPrice:  decimal.NewFromFloat(300.400),
	}
	tests := []struct {
		name         string
		exchangeMock exchangeBehaivor
		req          *pbexchange.GetRatesRequest
		want         *pbexchange.GetRatesResponse
		wantErr      bool
	}{
		{
			name: "test ok",
			exchangeMock: func(m *exchangemock.MockExchanger) {
				m.EXPECT().GetAndSaveRates(ctx, market).Return(rates, nil)
			},
			req: &pbexchange.GetRatesRequest{
				Market: pbexchange.Market_usdtrub,
			},
			want: &pbexchange.GetRatesResponse{
				Timestamp: rates.Timestamp.Unix(),
				Ask:       rates.AskPrice.InexactFloat64(),
				Bid:       rates.BidPrice.InexactFloat64(),
			},
			wantErr: false,
		},
		{
			name: "test error",
			exchangeMock: func(m *exchangemock.MockExchanger) {
				m.EXPECT().GetAndSaveRates(ctx, market).Return(domain.Rates{}, errors.New("error get rates"))
			},
			req: &pbexchange.GetRatesRequest{
				Market: pbexchange.Market_usdtrub,
			},
			want:    &pbexchange.GetRatesResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()
			mockExchanger := exchangemock.NewMockExchanger(c)

			e := NewExchange(logger, mockExchanger)
			tt.exchangeMock(mockExchanger)

			got, err := e.GetRates(ctx, tt.req)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}
