package handler

import (
	"currency_service/currency/internal/service"
	pb "currency_service/pkg/generated/currency"
	"log"
)

type CurrencyServer struct {
	pb.UnimplementedCurrencyServer
	service *service.CurrencyService
	logger  *log.Logger
}

func NewCurrencyServer(service *service.CurrencyService, logger *log.Logger) *CurrencyServer {
	return &CurrencyServer{
		service: service,
		logger:  logger,
	}
}
