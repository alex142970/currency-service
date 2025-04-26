package handler

import (
	"context"
	"currency_service/currency/internal/service"
	pb "currency_service/pkg/generated/currency"
	"google.golang.org/grpc"
	"time"
)

type Server struct {
	service *service.CurrencyService
	pb.UnimplementedCurrencyServer
}

func (s *Server) ExchangeRate(ctx context.Context, request *pb.ExchangeRateRequest) (*pb.ExchangeRateResponse, error) {
	date, _ := time.Parse("2006-01-02", request.Date)

	rate, err := s.service.GetRate(date)

	if err != nil {
		return nil, err
	}

	return &pb.ExchangeRateResponse{
		Rate: float32(*rate),
	}, nil
}

func (s *Server) ExchangeRateHistory(ctx context.Context, request *pb.ExchangeRateHistoryRequest) (*pb.ExchangeRateHistoryResponse, error) {
	from, _ := time.Parse("2006-01-02", request.DateFrom)
	to, _ := time.Parse("2006-01-02", request.DateTo)

	rates, err := s.service.GetRates(from, to)

	if err != nil {
		return nil, err
	}

	var result []float32

	for _, rate := range rates {
		result = append(result, float32(*rate))
	}

	return &pb.ExchangeRateHistoryResponse{
		Rate: result,
	}, nil
}

func NewServer(s *service.CurrencyService) *grpc.Server {
	serv := grpc.NewServer()

	pb.RegisterCurrencyServer(serv, &Server{
		service: s,
	})

	return serv
}
