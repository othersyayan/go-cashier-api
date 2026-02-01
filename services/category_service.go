package services

import (
	"go-cashier-api/models"
	"go-cashier-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetCategoryByID(id string) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}
