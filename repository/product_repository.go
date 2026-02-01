package repository

import (
	"database/sql"
	"errors"
	"kasir-api/models"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	rows, err := repo.db.Query("SELECT id, name, price, stock FROM products")
	if err != nil {
		log.Println("error", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("rows", rows)

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		log.Println("product", product)
		if err != nil {
			log.Println("error", err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	var product models.Product
	// Explicitly select columns to match Scan arguments
	row := repo.db.QueryRow("SELECT id, name, price, stock FROM products WHERE id = $1", id)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("Product not found")
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) CreateProduct(product models.Product) (*models.Product, error) {
	err := repo.db.QueryRow("INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id", product.Name, product.Price, product.Stock).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) UpdateProduct(product models.Product) (*models.Product, error) {
	row, err := repo.db.Exec("UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4", product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("Product not found")
	}
	return &product, nil
}

func (repo *ProductRepository) DeleteProduct(id int) error {
	row, err := repo.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Product not found")
	}

	return nil
}
