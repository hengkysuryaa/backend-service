package order

import (
	"context"
	"log"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

func (u *order) GetAll(ctx context.Context) ([]entity.Order, error) {
	orders, err := u.webRepo.GetOrders(ctx)
	if err != nil {
		log.Println(err)
		return []entity.Order{}, err
	}

	return orders, nil
}
