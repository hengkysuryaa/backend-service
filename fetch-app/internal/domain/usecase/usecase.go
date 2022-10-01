package usecase

import (
	"context"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
)

type Order interface {
	GetAll(ctx context.Context) ([]dto.Order, error)
	GetSummary(ctx context.Context, filter dto.GetSummaryFilter) (dto.OrderSummary, error)
}

type CurrencyConverter interface {
	ConvertCurrency(ctx context.Context, data dto.ConvertCurrencyRequest) (dto.ConvertCurrencyResponse, error)
}
