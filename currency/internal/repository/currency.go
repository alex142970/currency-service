package repository

import (
	"currency_service/currency/internal/dto"
	"database/sql"
	"time"
)

type ExchangeRateRepository interface {
	GetItemByDate(date time.Time) (*dto.ExchangeRate, error)
	GetItemsByDateRange(from time.Time, to time.Time) (*[]dto.ExchangeRate, error)
	Save(item *dto.ExchangeRate) error
}

type CurrencyPostgresRepository struct {
	db *sql.DB
}

func NewCurrencyPostgresRepository(db *sql.DB) *CurrencyPostgresRepository {
	return &CurrencyPostgresRepository{db: db}
}

func (r CurrencyPostgresRepository) GetItemByDate(date time.Time) (*dto.ExchangeRate, error) {
	row := r.db.QueryRow("SELECT id, base_currency, target_currency, rate, timestamp FROM exchange_rates WHERE date(timestamp) = $1 LIMIT 1", date.Format("2006-01-02"))

	var item dto.ExchangeRate

	err := row.Scan(&item.Id, &item.BaseCurrency, &item.TargetCurrency, &item.Rate, &item.Timestamp)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r CurrencyPostgresRepository) GetItemsByDateRange(from time.Time, to time.Time) (*[]dto.ExchangeRate, error) {
	rows, err := r.db.Query("SELECT id, base_currency, target_currency, rate, timestamp FROM exchange_rates WHERE date(timestamp) >= $1 AND date(timestamp) <= $2", from.Format("2006-01-02"), to.Format("2006-01-02"))

	var items []dto.ExchangeRate

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var item dto.ExchangeRate

		err := rows.Scan(&item.Id, &item.BaseCurrency, &item.TargetCurrency, &item.Rate, &item.Timestamp)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func (r CurrencyPostgresRepository) Save(item *dto.ExchangeRate) error {

	_, err := r.db.Exec("INSERT INTO exchange_rates (base_currency, target_currency, rate) VALUES ($1, $2, $3)", item.BaseCurrency, item.TargetCurrency, item.Rate)

	if err != nil {
		return err
	}

	return nil
}
