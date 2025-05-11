package repository

import (
	"context"
	"currency_service/currency/internal/dto"
	"database/sql"
	"fmt"
	"time"
)

type CurrencyPostgresRepository struct {
	db *sql.DB
}

func NewCurrencyPostgresRepository(db *sql.DB) *CurrencyPostgresRepository {
	return &CurrencyPostgresRepository{db: db}
}

func (r CurrencyPostgresRepository) GetItemByDate(ctx context.Context, date time.Time) (*dto.ExchangeRate, error) {
	query := `SELECT id, base_currency, target_currency, rate, timestamp FROM exchange_rates WHERE date(timestamp) = $1 LIMIT 1`

	row := r.db.QueryRowContext(ctx, query, date.Format("2006-01-02"))

	if row == nil {
		return nil, fmt.Errorf("no exchange rate found for %s", date.Format("2006-01-02"))
	}

	var item dto.ExchangeRate

	err := row.Scan(&item.Id, &item.BaseCurrency, &item.TargetCurrency, &item.Rate, &item.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("error scan row %s: %w", date.Format("2006-01-02"), err)
	}

	return &item, nil
}

func (r CurrencyPostgresRepository) GetItemsByDateRange(ctx context.Context, from time.Time, to time.Time) ([]dto.ExchangeRate, error) {
	query := `
		SELECT id, base_currency, target_currency, rate, timestamp 
		FROM exchange_rates WHERE date(timestamp) >= $1 AND date(timestamp) <= $2
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		from.Format("2006-01-02"),
		to.Format("2006-01-02"),
	)

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to get items by date range: %w", err)
	}

	var items []dto.ExchangeRate

	for rows.Next() {
		var item dto.ExchangeRate

		err := rows.Scan(&item.Id, &item.BaseCurrency, &item.TargetCurrency, &item.Rate, &item.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while rows iteration: %w", err)
	}

	return items, nil
}

func (r CurrencyPostgresRepository) Save(ctx context.Context, item *dto.ExchangeRate) error {

	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO exchange_rates (base_currency, target_currency, rate) VALUES ($1, $2, $3)",
		item.BaseCurrency, item.TargetCurrency, item.Rate,
	)

	if err != nil {
		return fmt.Errorf("could not save exchange rate: %w", err)
	}

	return nil
}
