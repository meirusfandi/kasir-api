package services

import (
	"kasir-api/models"
	"kasir-api/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts(name string) ([]models.Product, error) {
	return s.repo.GetAllProducts(name)
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(product models.Product) (*models.Product, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product models.Product) (*models.Product, error) {
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
