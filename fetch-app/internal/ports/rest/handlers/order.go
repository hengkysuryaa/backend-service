package handlers

import (
	"net/http"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/usecase"
	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/helpers"
)

type OrderHandlers struct {
	orderUsecase usecase.Order
}

func NewOrderHandler(orderUsecase usecase.Order) *OrderHandlers {
	return &OrderHandlers{
		orderUsecase: orderUsecase,
	}
}

func (h *OrderHandlers) GetAll(rw http.ResponseWriter, r *http.Request) {
	orders, err := h.orderUsecase.GetAll(r.Context())
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data, err := helpers.BuildJSON(rw, orders)
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
