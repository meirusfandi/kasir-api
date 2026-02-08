package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-api/models"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAllProducts(name string) ([]models.Product, error) {
	var products []models.Product
	query := `
		SELECT id, name, price, stock
		FROM products
	`
	args := []interface{}{}
	if name != "" {
		query += `WHERE name LIKE $1`
		args = append(args, "%"+name+"%")
	}

	fmt.Println("query", query)
	fmt.Println("args", args)

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		log.Println("error", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("rows", rows)

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID, &product.Name, &product.Price, &product.Stock,
		)
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
	product.Category = &models.Category{} // Initialize pointer
	query := `
		SELECT p.id, p.name, p.price, p.stock, c.id, c.name 
		FROM products p
		JOIN category c ON p.category_id = c.id
		WHERE p.id = $1
	`
	row := repo.db.QueryRow(query, id)
	err := row.Scan(
		&product.ID, &product.Name, &product.Price, &product.Stock,
		&product.Category.ID, &product.Category.Name,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("Product not found")
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) CreateProduct(product models.Product) (*models.Product, error) {
	var categoryID int
	if product.Category != nil {
		categoryID = product.Category.ID
	}
	err := repo.db.QueryRow("INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		product.Name, product.Price, product.Stock, categoryID).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) UpdateProduct(product models.Product) (*models.Product, error) {
	var categoryID int
	if product.Category != nil {
		categoryID = product.Category.ID
	}
	row, err := repo.db.Exec("UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5",
		product.Name, product.Price, product.Stock, categoryID, product.ID)
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
