package handlers

import (
	"encoding/json"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/usecase"
	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/helpers"
)

type CurrencyConverterHandlers struct {
	currencyConverterUsecase usecase.CurrencyConverter
}

func NewCurrencyConverterHandler(currencyConverterUsecase usecase.CurrencyConverter) *CurrencyConverterHandlers {
	return &CurrencyConverterHandlers{
		currencyConverterUsecase: currencyConverterUsecase,
	}
}

func (h *CurrencyConverterHandlers) ConvertCurrency(rw http.ResponseWriter, r *http.Request) {
	var body dto.ConvertCurrencyRequest
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

	conversionResult, err := h.currencyConverterUsecase.ConvertCurrency(r.Context(), body)
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data, err := helpers.BuildJSON(rw, conversionResult)
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
