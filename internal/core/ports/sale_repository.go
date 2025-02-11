package ports

import "github.com/alemelomeza/effective-succotash/internal/core/domain/entities"

type SaleRepository interface {
	GetProductsByIDs(ids []int) ([]entities.Product, error)
}
