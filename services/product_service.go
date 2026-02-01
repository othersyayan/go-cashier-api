package services

import (
	"go-cashier-api/models"
	"go-cashier-api/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repository.Create(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repository.GetAll()
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repository.FindByID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repository.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repository.Delete(id)
}
