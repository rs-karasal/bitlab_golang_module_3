// internal/repository/product_repo.go
package repository

import (
	"products_service_with_dto/internal/entities"
)

type ProductRepository struct {
	data   map[int]entities.Product
	nextID int
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		data:   map[int]entities.Product{},
		nextID: 1,
	}
}

func (r *ProductRepository) Create(product entities.Product) (entities.Product, error) {
	r.data[r.nextID] = product
	r.nextID++

	return product, nil
}

func (r *ProductRepository) List() []entities.Product {
	products := make([]entities.Product, 0, len(r.data))
	for _, product := range r.data {
		products = append(products, product)
	}
	return products
}
