package repository

import (
	"context"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

type WebRepo interface {
	GetOrders(ctx context.Context) ([]entity.Order, error)
	ConvertCurrency(ctx context.Context, data dto.ConvertCurrencyRequest) (entity.CurrencyConverter, error)
}
