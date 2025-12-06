// internal/handlers/dto/product_dto.go
package dto

import "app_with_postgres/internal/entities"

type CreateProductRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ProductResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// NewProductFromCreateRequest создаёт новый продукт из запроса на создание
func NewProductFromCreateRequest(req CreateProductRequest) entities.Product {
	return entities.Product{
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	}
}

// NewProductResponse создаёт новый ответ на запрос продукта
func NewProductResponse(p *entities.Product) ProductResponse {
	return ProductResponse{
		ID:       p.ID,
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}
}
