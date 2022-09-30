package order

import (
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/repository"
)

type order struct {
	webRepo repository.WebRepo
}

func New(webRepo repository.WebRepo) *order {
	return &order{
		webRepo: webRepo,
	}
}
