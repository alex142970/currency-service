package grpc_client

import (
	"context"
	"currency_service/pkg/generated/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type CurrencyClient struct {
	api com_currency.CurrencyClient
}

func NewCurrencyClient(ctx context.Context, addr string) (*CurrencyClient, error) {
	connection, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return &CurrencyClient{
		api: com_currency.NewCurrencyClient(connection),
	}, nil
}

func (cc *CurrencyClient) ExchangeRate(ctx context.Context, date time.Time) (*com_currency.ExchangeRateResponse, error) {
	response, err := cc.api.ExchangeRate(ctx, &com_currency.ExchangeRateRequest{
		Date: date.Format("2006-01-02"),
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (cc *CurrencyClient) ExchangeRateHistory(ctx context.Context, from time.Time, to time.Time) (*com_currency.ExchangeRateHistoryResponse, error) {
	response, err := cc.api.ExchangeRateHistory(ctx, &com_currency.ExchangeRateHistoryRequest{
		DateFrom: from.Format("2006-01-02"),
		DateTo:   to.Format("2006-01-02"),
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
