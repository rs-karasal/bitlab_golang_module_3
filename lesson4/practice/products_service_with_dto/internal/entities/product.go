// internal/entities/product.go
package entities

type Product struct {
	// TODO: Описать сущность Product (товар) с полями: ID, Name, Price, Quantity.
}

// Validate проверяет, что товар имеет все необходимые поля
func (p Product) Validate() error {
	// TODO: - Написать метод Validate():
	// - Name не пустой;
	// - Price >= 0;
	// - Quantity >= 0.

	return nil
}
