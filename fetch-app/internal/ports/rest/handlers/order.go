package handlers

import (
	"encoding/json"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/usecase"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/ports/rest/handlers/request"
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

func (h *OrderHandlers) GetSummary(rw http.ResponseWriter, r *http.Request) {
	var body request.GetSummaryRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	summary, err := h.orderUsecase.GetSummary(r.Context(), dto.GetSummaryFilter{
		AreaProvinsi: body.AreaProvinsi,
		Tanggal:      body.Tanggal,
	})
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data, err := helpers.BuildJSON(rw, summary)
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
