// internal/repository/postgres_product_repo.go
package repository

import (
	"database/sql"
	"errors"

	"app_with_postgres/internal/entities"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Create(p entities.Product) (entities.Product, error) {
	query := `
        INSERT INTO products (name, price, quantity)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	err := r.db.QueryRow(query, p.Name, p.Price, p.Quantity).Scan(&p.ID)
	if err != nil {
		return entities.Product{}, err
	}

	return p, nil
}

func (r *PostgresProductRepository) GetByID(id int) (entities.Product, error) {
	query := `
        SELECT id, name, price, quantity
        FROM products
        WHERE id = $1;
    `

	var p entities.Product
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Product{}, errors.New("product not found")
		}
		return entities.Product{}, err
	}

	return p, nil
}

func (r *PostgresProductRepository) List() ([]entities.Product, error) {
	query := `
        SELECT id, name, price, quantity
        FROM products
        ORDER BY id;
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var p entities.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
