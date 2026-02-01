package usecase

import (
	"go-cashier-api/internal/entity"
	"go-cashier-api/internal/repository"
)

type ProductUsecase interface {
	GetAll() ([]entity.Product, error)
	GetByID(id int) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	Update(id int, product entity.Product) (entity.Product, error)
	Delete(id int) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

func (u *productUsecase) GetAll() ([]entity.Product, error) {
	return u.repo.FindAll()
}

func (u *productUsecase) GetByID(id int) (entity.Product, error) {
	return u.repo.FindByID(id)
}

func (u *productUsecase) Create(product entity.Product) (entity.Product, error) {
	return u.repo.Create(product)
}

func (u *productUsecase) Update(id int, product entity.Product) (entity.Product, error) {
	return u.repo.Update(id, product)
}

func (u *productUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
