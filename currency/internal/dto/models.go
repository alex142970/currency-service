package dto

import "time"

type ExchangeRate struct {
	Id             uint64
	BaseCurrency   string
	TargetCurrency string
	Rate           float64
	Timestamp      time.Time
}
