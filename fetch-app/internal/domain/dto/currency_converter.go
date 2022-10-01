package dto

type ConvertCurrencyRequest struct {
	From   string  `json:"from" validate:"required,oneof=IDR USD"`
	To     string  `json:"to" validate:"required,oneof=IDR USD"`
	Amount float64 `json:"amount" validate:"required"`
}

type CurrencyConversionCache struct {
	Rate          float64
	FetchAtUnix   int64
	ExpiredAtUnix int64
}

type ConvertCurrencyResponse struct {
	Rate          float64                `json:"rate"`
	RateTimestamp int64                  `json:"rate_timestamp"`
	Request       ConvertCurrencyRequest `json:"request"`
	Result        float64                `json:"result"`
}
