package service

import (
	"currency_service/currency/internal/app"
	"currency_service/currency/internal/config"
	"currency_service/currency/internal/dto"
	"currency_service/currency/internal/repository"
	"time"
)

type CurrencyService struct {
	Repo repository.ExchangeRateRepository
	Cfg  *config.Config
}

func NewCurrencyService(app *app.App) *CurrencyService {
	return &CurrencyService{
		Repo: repository.NewCurrencyPostgresRepository(app.Db),
		Cfg:  app.Config,
	}
}

func (s CurrencyService) AddRate(rate float64, time time.Time) error {
	err := s.Repo.Save(&dto.ExchangeRate{
		BaseCurrency:   s.Cfg.Currency.Base,
		TargetCurrency: s.Cfg.Currency.Target,
		Rate:           rate,
		Timestamp:      time,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s CurrencyService) GetRate(date time.Time) (*float64, error) {
	item, err := s.Repo.GetItemByDate(date)

	if err != nil {
		return nil, err
	}

	return &item.Rate, nil
}

func (s CurrencyService) GetRates(from time.Time, to time.Time) ([]*float64, error) {
	items, err := s.Repo.GetItemsByDateRange(from, to)

	if err != nil {
		return nil, err
	}

	var rates []*float64

	for _, item := range *items {
		rates = append(rates, &item.Rate)
	}

	return rates, nil
}
