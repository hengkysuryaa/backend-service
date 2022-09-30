package usecase

import (
	"context"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

type Order interface {
	GetAll(ctx context.Context) ([]entity.Order, error)
}
