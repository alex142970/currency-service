package handler

import (
	"context"
	pb "currency_service/pkg/generated/currency"
	"fmt"
	"time"
)

func (s *CurrencyServer) ExchangeRate(ctx context.Context, request *pb.ExchangeRateRequest) (*pb.ExchangeRateResponse, error) {
	op := "ExchangeRate"

	s.logger.SetPrefix(op)

	date, _ := time.Parse("2006-01-02", request.Date)

	rate, err := s.service.GetRate(ctx, date)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &pb.ExchangeRateResponse{
		Rate: float32(*rate),
	}, nil
}

func (s *CurrencyServer) ExchangeRateHistory(ctx context.Context, request *pb.ExchangeRateHistoryRequest) (*pb.ExchangeRateHistoryResponse, error) {
	op := "ExchangeRateHistory"

	s.logger.SetPrefix(op)

	from, _ := time.Parse("2006-01-02", request.DateFrom)
	to, _ := time.Parse("2006-01-02", request.DateTo)

	rates, err := s.service.GetRates(ctx, from, to)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var result []float32

	for _, rate := range rates {
		result = append(result, float32(*rate))
	}

	return &pb.ExchangeRateHistoryResponse{
		Rate: result,
	}, nil
}
