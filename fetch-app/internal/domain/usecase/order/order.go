package order

import (
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/repository"
	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/cache"
)

type order struct {
	webRepo repository.WebRepo
	cache   *cache.MapCache
}

func New(webRepo repository.WebRepo, cache *cache.MapCache) *order {
	return &order{
		webRepo: webRepo,
		cache:   cache,
	}
}
