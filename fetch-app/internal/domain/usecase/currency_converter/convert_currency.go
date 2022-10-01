package currency_converter

import (
	"context"
	"log"
	"time"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

func (u *currencyConverter) ConvertCurrency(ctx context.Context, data dto.ConvertCurrencyRequest) (dto.ConvertCurrencyResponse, error) {
	// first, get value of currency conversion from cache
	val := u.cache.Get(entity.CURRENCY_CONVERSION_BASE_KEY + data.From + data.To)
	if val != nil {
		// check cache expired time
		currencyConversion := val.(dto.CurrencyConversionCache)
		if currencyConversion.ExpiredAtUnix >= time.Now().UTC().Unix() {
			return dto.ConvertCurrencyResponse{
				Rate:          currencyConversion.Rate,
				RateTimestamp: currencyConversion.FetchAtUnix,
				Request:       data,
				Result:        currencyConversion.Rate * data.Amount,
			}, nil
		}
	}

	// if value is nil or value is expired, then fetch to repository
	currencyConversion, err := u.webRepo.ConvertCurrency(ctx, data)
	if err != nil {
		log.Println(err)
		return dto.ConvertCurrencyResponse{}, err
	}

	// set value to MapCache
	cacheVal := dto.CurrencyConversionCache{
		Rate:          currencyConversion.Info.Rate,
		FetchAtUnix:   int64(currencyConversion.Info.Timestamp),
		ExpiredAtUnix: time.Now().UTC().Add(time.Hour * time.Duration(entity.CACHE_EXPIRED_IN_HOUR)).Unix(),
	}
	u.cache.Store(entity.CURRENCY_CONVERSION_BASE_KEY+data.From+data.To, cacheVal)

	return dto.ConvertCurrencyResponse{
		Rate:          currencyConversion.Info.Rate,
		RateTimestamp: int64(currencyConversion.Info.Timestamp),
		Request:       data,
		Result:        currencyConversion.Result,
	}, nil
}
