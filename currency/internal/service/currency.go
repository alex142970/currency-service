package service

import (
	"context"
	"currency_service/currency/internal/config"
	"currency_service/currency/internal/dto"
	"fmt"
	"log"
	"time"
)

type ExchangeRateRepository interface {
	GetItemByDate(ctx context.Context, date time.Time) (*dto.ExchangeRate, error)
	GetItemsByDateRange(ctx context.Context, from time.Time, to time.Time) ([]dto.ExchangeRate, error)
	Save(ctx context.Context, item *dto.ExchangeRate) error
}

type CurrencyService struct {
	Repo   ExchangeRateRepository
	Cfg    *config.Config
	Logger *log.Logger
}

func NewCurrencyService(repo ExchangeRateRepository, cfg *config.Config, logger *log.Logger) *CurrencyService {
	return &CurrencyService{
		Repo:   repo,
		Cfg:    cfg,
		Logger: logger,
	}
}

func (s CurrencyService) AddRate(ctx context.Context, rate float64, time time.Time) error {
	err := s.Repo.Save(ctx, &dto.ExchangeRate{
		BaseCurrency:   s.Cfg.Currency.Base,
		TargetCurrency: s.Cfg.Currency.Target,
		Rate:           rate,
		Timestamp:      time,
	})

	if err != nil {
		return fmt.Errorf("failed to add rate: %w", err)
	}

	return nil
}

func (s CurrencyService) GetRate(ctx context.Context, date time.Time) (*float64, error) {
	item, err := s.Repo.GetItemByDate(ctx, date)

	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	return &item.Rate, nil
}

func (s CurrencyService) GetRates(ctx context.Context, from time.Time, to time.Time) ([]*float64, error) {
	items, err := s.Repo.GetItemsByDateRange(ctx, from, to)

	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}

	var rates []*float64

	for _, item := range items {
		rates = append(rates, &item.Rate)
	}

	return rates, nil
}
