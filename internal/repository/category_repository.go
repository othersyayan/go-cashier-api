package repository

import (
	"go-cashier-api/internal/entity"
)

type CategoryRepository interface {
	FindAll() ([]entity.Category, error)
	FindByID(id int) (entity.Category, error)
	Create(category entity.Category) (entity.Category, error)
	Update(id int, category entity.Category) (entity.Category, error)
	Delete(id int) error
}

type inMemoryCategoryRepository struct {
	categories []entity.Category
}

func NewInMemoryCategoryRepository() CategoryRepository {
	return &inMemoryCategoryRepository{
		categories: []entity.Category{},
	}
}

func (r *inMemoryCategoryRepository) FindAll() ([]entity.Category, error) {
	return r.categories, nil
}

func (r *inMemoryCategoryRepository) FindByID(id int) (entity.Category, error) {
	for _, c := range r.categories {
		if c.ID == id {
			return c, nil
		}
	}
	return entity.Category{}, ErrNotFound
}

func (r *inMemoryCategoryRepository) Create(category entity.Category) (entity.Category, error) {
	category.ID = len(r.categories) + 1
	r.categories = append(r.categories, category)
	return category, nil
}

func (r *inMemoryCategoryRepository) Update(id int, category entity.Category) (entity.Category, error) {
	for i, c := range r.categories {
		if c.ID == id {
			category.ID = id
			r.categories[i] = category
			return category, nil
		}
	}
	return entity.Category{}, ErrNotFound
}

func (r *inMemoryCategoryRepository) Delete(id int) error {
	for i, c := range r.categories {
		if c.ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
