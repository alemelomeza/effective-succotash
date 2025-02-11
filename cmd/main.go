package main

import (
	"net/http"

	"github.com/alemelomeza/effective-succotash/internal/infrastructure/api"
)

func main() {
	h := api.CreateSaleHandler{
		SaleRepository: nil,
	}
	http.HandleFunc("POST /products", h.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
