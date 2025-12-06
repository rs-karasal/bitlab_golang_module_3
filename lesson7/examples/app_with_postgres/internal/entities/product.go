// internal/entities/product.go
package entities

import "errors"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func (p Product) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Price == 0 {
		return errors.New("price is required")
	}

	if p.Quantity == 0 {
		return errors.New("quantity is required")
	}

	return nil
}
