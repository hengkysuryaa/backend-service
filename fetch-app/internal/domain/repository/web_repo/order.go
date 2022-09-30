package web_repo

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
)

func (r *webrepo) GetOrders(ctx context.Context) ([]entity.Order, error) {
	var orders []entity.Order

	req, err := http.NewRequest("GET", r.resourceUrl, nil)
	if err != nil {
		return []entity.Order{}, err
	}

	response, err := r.httpClient.Do(req)
	if err != nil {
		return []entity.Order{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []entity.Order{}, err
	}

	err = json.Unmarshal(body, &orders)
	if err != nil {
		return []entity.Order{}, err
	}

	return orders, nil
}
