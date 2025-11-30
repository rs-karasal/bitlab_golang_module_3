// internal/service/product_service.go
package service

import "products_service_with_dto/internal/entities"

type ProductRepository interface {
	Create(product entities.Product) (entities.Product, error)
	List() []entities.Product
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(r ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) CreateProduct(p entities.Product) (entities.Product, error) {
	if err := p.Validate(); err != nil {
		return entities.Product{}, err
	}

	return s.repo.Create(p)
}

func (s *ProductService) ListProducts() []entities.Product {
	return s.repo.List()
}
