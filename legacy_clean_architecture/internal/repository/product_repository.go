package repository

import (
	"errors"
	"go-cashier-api/internal/entity"
)

var (
	ErrNotFound = errors.New("product not found")
)

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindByID(id int) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	Update(id int, product entity.Product) (entity.Product, error)
	Delete(id int) error
}

type inMemoryProductRepository struct {
	products []entity.Product
}

func NewInMemoryProductRepository() ProductRepository {
	return &inMemoryProductRepository{
		products: []entity.Product{
			{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
			{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
			{ID: 3, Name: "kecap", Price: 12000, Stock: 20},
		},
	}
}

func (r *inMemoryProductRepository) FindAll() ([]entity.Product, error) {
	return r.products, nil
}

func (r *inMemoryProductRepository) FindByID(id int) (entity.Product, error) {
	for _, p := range r.products {
		if p.ID == id {
			return p, nil
		}
	}
	return entity.Product{}, ErrNotFound
}

func (r *inMemoryProductRepository) Create(product entity.Product) (entity.Product, error) {
	product.ID = len(r.products) + 1
	r.products = append(r.products, product)
	return product, nil
}

func (r *inMemoryProductRepository) Update(id int, product entity.Product) (entity.Product, error) {
	for i, p := range r.products {
		if p.ID == id {
			product.ID = id
			r.products[i] = product
			return product, nil
		}
	}
	return entity.Product{}, ErrNotFound
}

func (r *inMemoryProductRepository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
