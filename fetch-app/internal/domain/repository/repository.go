package repository

import (
	"context"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

type WebRepo interface {
	GetOrders(ctx context.Context) ([]entity.Order, error)
}
