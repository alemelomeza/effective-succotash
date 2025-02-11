package api

import (
	"encoding/json"
	"net/http"

	"github.com/alemelomeza/effective-succotash/internal/core/domain/usecases"
	"github.com/alemelomeza/effective-succotash/internal/core/ports"
)

type CreateSaleRequest struct {
	Products []int `json:"products"`
}
type CreateSaleResponse struct {
	Total float64 `json:"total"`
}
type CreateSaleHandler struct {
	SaleRepository ports.SaleRepository
}

func (h *CreateSaleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request CreateSaleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	products, err := h.SaleRepository.GetProductsByIDs(request.Products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	total := usecases.CalculateTotal(products, nil)
	response := CreateSaleResponse{
		Total: total,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
