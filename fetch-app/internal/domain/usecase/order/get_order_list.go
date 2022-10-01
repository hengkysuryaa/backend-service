package order

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

func (u *order) GetAll(ctx context.Context) ([]dto.Order, error) {
	// fetch currency conversion
	currencyConversion, err := u.fetchCurrencyConversion(ctx, "IDR", "USD", 1)
	if err != nil {
		log.Println(err)
		return []dto.Order{}, err
	}

	orders, err := u.webRepo.GetOrders(ctx)
	if err != nil {
		log.Println(err)
		return []dto.Order{}, err
	}

	// iterate orders data
	var ordersList []dto.Order
	for _, order := range orders {
		orderDto := dto.Order{
			UUID:         order.UUID,
			Komoditas:    order.Komoditas,
			AreaProvinsi: order.AreaProvinsi,
			AreaKota:     order.AreaKota,
			Size:         order.Size,
		}

		if order.Price != nil {
			idrPrice, _ := strconv.ParseFloat(*order.Price, 64)
			usdPrice := currencyConversion.Rate * idrPrice

			orderDto.USDPrice = &usdPrice
			orderDto.IDRPrice = &idrPrice
		}

		ordersList = append(ordersList, orderDto)
	}

	return ordersList, nil
}

func (u *order) fetchCurrencyConversion(ctx context.Context, from, to string, amount float64) (dto.CurrencyConversionCache, error) {
	var currencyConversion dto.CurrencyConversionCache
	val := u.cache.Get(entity.CURRENCY_CONVERSION_BASE_KEY + from + to)
	if val != nil {
		// check cache expired time
		currencyConversion = val.(dto.CurrencyConversionCache)
		if currencyConversion.ExpiredAtUnix >= time.Now().UTC().Unix() {
			return currencyConversion, nil
		}
	}

	// if value is nil or value is expired, then fetch to repository
	convertResult, err := u.webRepo.ConvertCurrency(ctx, dto.ConvertCurrencyRequest{
		From:   from,
		To:     to,
		Amount: amount,
	})
	if err != nil {
		return dto.CurrencyConversionCache{}, err
	}

	// set value to MapCache
	cacheVal := dto.CurrencyConversionCache{
		Rate:          convertResult.Info.Rate,
		FetchAtUnix:   int64(convertResult.Info.Timestamp),
		ExpiredAtUnix: time.Now().UTC().Add(time.Hour * time.Duration(entity.CACHE_EXPIRED_IN_HOUR)).Unix(),
	}
	u.cache.Store(entity.CURRENCY_CONVERSION_BASE_KEY+from+to, cacheVal)

	return cacheVal, nil
}
