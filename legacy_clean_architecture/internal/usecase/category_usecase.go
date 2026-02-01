package usecase

import (
	"go-cashier-api/internal/entity"
	"go-cashier-api/internal/repository"
)

type CategoryUsecase interface {
	GetAll() ([]entity.Category, error)
	GetByID(id int) (entity.Category, error)
	Create(category entity.Category) (entity.Category, error)
	Update(id int, category entity.Category) (entity.Category, error)
	Delete(id int) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		repo: repo,
	}
}

func (u *categoryUsecase) GetAll() ([]entity.Category, error) {
	return u.repo.FindAll()
}

func (u *categoryUsecase) GetByID(id int) (entity.Category, error) {
	return u.repo.FindByID(id)
}

func (u *categoryUsecase) Create(category entity.Category) (entity.Category, error) {
	return u.repo.Create(category)
}

func (u *categoryUsecase) Update(id int, category entity.Category) (entity.Category, error) {
	return u.repo.Update(id, category)
}

func (u *categoryUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
