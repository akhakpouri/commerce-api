package category

import (
	"commerce/api/internal/dto/category"

	"gorm.io/gorm"
)

type CategoryServiceI interface {
	GetById(id int) (*category.Category, error)
	GetAll() ([]*category.Category, error)
	Create(category *category.Category) (*category.Category, error)
	Update(id int, category *category.Category) (*category.Category, error)
	Delete(id int) error
}

type CategoryService struct {
	db *gorm.DB
}
