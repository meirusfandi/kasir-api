package services

import (
	"kasir-api/models"
	"kasir-api/repository"
)

type CategoryService struct {
	repository *repository.CategoryRepository
}

func NewCategoryService(repository *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repository: repository}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repository.GetAllCategories()
}

func (s *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	return s.repository.GetCategoryByID(id)
}

func (s *CategoryService) CreateCategory(category models.Category) (*models.Category, error) {
	return s.repository.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(category models.Category) (*models.Category, error) {
	return s.repository.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.repository.DeleteCategory(id)
}
