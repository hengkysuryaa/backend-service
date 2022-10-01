package usecase

import (
	"context"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

type Order interface {
	GetAll(ctx context.Context) ([]entity.Order, error)
	GetSummary(ctx context.Context, filter dto.GetSummaryFilter) (dto.OrderSummary, error)
}

type CurrencyConverter interface {
	ConvertCurrency(ctx context.Context, data dto.ConvertCurrencyRequest) (dto.ConvertCurrencyResponse, error)
}
