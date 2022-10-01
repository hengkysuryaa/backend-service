package currency_converter

import (
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/repository"
	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/cache"
)

type currencyConverter struct {
	webRepo repository.WebRepo
	cache   *cache.MapCache
}

func New(webRepo repository.WebRepo, cache *cache.MapCache) *currencyConverter {
	return &currencyConverter{
		webRepo: webRepo,
		cache:   cache,
	}
}
