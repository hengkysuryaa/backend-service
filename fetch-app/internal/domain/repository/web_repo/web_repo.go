package web_repo

import "net/http"

type webrepo struct {
	httpClient              *http.Client
	resourceUrl             string
	currencyConverterUrl    string
	currencyConverterAPIKey string
}

func New(httpClient *http.Client,
	resourceUrl string,
	currencyConverterUrl string,
	currencyConverterAPIKey string,
) *webrepo {
	return &webrepo{
		httpClient:              httpClient,
		resourceUrl:             resourceUrl,
		currencyConverterUrl:    currencyConverterUrl,
		currencyConverterAPIKey: currencyConverterAPIKey,
	}
}
