package repository

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	query := `
		SELECT c.id, c.name, c.description, COUNT(p.id) as product_count
		FROM category c
		LEFT JOIN products p ON c.id = p.category_id
		GROUP BY c.id, c.name, c.description
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.ProductCount)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	var category models.Category
	query := `
		SELECT c.id, c.name, c.description, COUNT(p.id) as product_count
		FROM category c
		LEFT JOIN products p ON c.id = p.category_id
		WHERE c.id = $1
		GROUP BY c.id, c.name, c.description
	`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&category.ID, &category.Name, &category.Description, &category.ProductCount)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) CreateCategory(category models.Category) (*models.Category, error) {
	err := r.db.QueryRow("INSERT INTO category (name, description) VALUES ($1, $2) RETURNING id", category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) UpdateCategory(category models.Category) (*models.Category, error) {
	row, err := r.db.Exec("UPDATE category SET name = $1, description = $2 WHERE id = $3", category.Name, category.Description, category.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("Category not found")
	}
	return &category, nil
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	row, err := r.db.Exec("DELETE FROM category WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Category not found")
	}

	return nil
}
