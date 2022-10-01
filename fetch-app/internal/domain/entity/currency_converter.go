package entity

var CURRENCY_CONVERSION_BASE_KEY = "currency-conversion:"
var CACHE_EXPIRED_IN_HOUR = 2

type CurrencyConverter struct {
	Date       string `json:"date"`
	Historical string `json:"historical"`
	Info       struct {
		Rate      float64 `json:"rate"`
		Timestamp int     `json:"timestamp"`
	} `json:"info"`
	Query struct {
		Amount int    `json:"amount"`
		From   string `json:"from"`
		To     string `json:"to"`
	} `json:"query"`
	Result  float64 `json:"result"`
	Success bool    `json:"success"`
}
