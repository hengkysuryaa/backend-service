package web_repo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

func (r *webrepo) ConvertCurrency(ctx context.Context, data dto.ConvertCurrencyRequest) (entity.CurrencyConverter, error) {
	var currencyConverter entity.CurrencyConverter

	fullURL := fmt.Sprintf("%s/convert?to=%s&from=%s&amount=%f", r.currencyConverterUrl, data.To, data.From, data.Amount)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return entity.CurrencyConverter{}, err
	}

	req.Header.Set("apikey", r.currencyConverterAPIKey)
	response, err := r.httpClient.Do(req)
	if err != nil {
		return entity.CurrencyConverter{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return entity.CurrencyConverter{}, err
	}

	err = json.Unmarshal(body, &currencyConverter)
	if err != nil {
		return entity.CurrencyConverter{}, err
	}

	return currencyConverter, nil
}
