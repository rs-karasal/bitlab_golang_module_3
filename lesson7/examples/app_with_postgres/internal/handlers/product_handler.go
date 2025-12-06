// internal/handlers/product_handler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"app_with_postgres/internal/entities"
	"app_with_postgres/internal/handlers/dto"
)

type ProductService interface {
	CreateProduct(p entities.Product) (entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	ListProducts() ([]entities.Product, error)
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(s ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) HandlerProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createProduct(w, r)
	case http.MethodGet:
		h.listProducts(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	product := dto.NewProductFromCreateRequest(req)
	created, err := h.service.CreateProduct(product)
	if err != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}
	createdResp := dto.NewProductResponse(&created)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdResp)
}

func (h *ProductHandler) listProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts()
	if err != nil {
		http.Error(w, "failed to list products", http.StatusInternalServerError)
		return
	}

	productsResp := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		productsResp[i] = dto.NewProductResponse(&product)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productsResp)
}
