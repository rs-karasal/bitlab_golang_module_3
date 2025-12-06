package service

import "app_with_postgres/internal/entities"

type ProductRepository interface {
	Create(p entities.Product) (entities.Product, error)
	GetByID(id int) (entities.Product, error)
	List() ([]entities.Product, error)
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

func (s *ProductService) GetProductByID(id int) (entities.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) ListProducts() ([]entities.Product, error) {
	return s.repo.List()
}
