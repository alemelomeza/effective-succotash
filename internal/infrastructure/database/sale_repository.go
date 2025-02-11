package database

import "github.com/alemelomeza/effective-succotash/internal/core/domain/entities"

type PostgresSaleRepository struct {
	// ... connection to Postgres data base
}

func (repo *PostgresSaleRepository) GetProductsByIDs(ids []int) ([]entities.Product, error) {
	// ... logic to get products form the data base
	return nil, nil
}
