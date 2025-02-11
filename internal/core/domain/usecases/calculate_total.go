package usecases

import "github.com/alemelomeza/effective-succotash/internal/core/domain/entities"

func CalculateTotal(products []entities.Product, productIDs []int) float64 {
	total := 0.0
	for _, productID := range productIDs {
		for _, product := range products {
			if product.ID == productID {
				total += product.Price
				break
			}
		}
	}
	return total
}
