// internal/handlers/dto/product_dto.go
package dto

import "products_service_with_dto/internal/entities"

type CreateProductRequest struct {
	// TODO: Описать структуру CreateProductRequest с полями: Name, Price, Quantity.
}

type ProductResponse struct {
	// TODO: Описать структуру ProductResponse с полями: ID, Name, Price, Quantity.
}

func NewProductFromCreateRequest(req CreateProductRequest) entities.Product {
	// TODO: - Написать функцию NewProductFromCreateRequest()

	return entities.Product{}
}

func NewProductResponse(p entities.Product) ProductResponse {
	// TODO: - Написать функцию NewProductResponse()

	return ProductResponse{}
}
