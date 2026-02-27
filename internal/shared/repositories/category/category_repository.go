package category

import (
	"commerce/internal/shared/models"

	"gorm.io/gorm"
)

type CategoryRepositoryI interface {
	GetById(id uint) (*models.Category, error)
	GetByParentId(parentId uint) ([]*models.Category, error)
	GetAll() ([]*models.Category, error)
	Save(category *models.Category) error
	Delete(id uint, hard bool) error
}

type CategoryRepository struct {
	db *gorm.DB
}

func (r *CategoryRepository) GetById(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category.Id, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetByParentId(parentId uint) ([]*models.Category, error) {
	var categories []*models.Category
	if err := r.db.Where("parent_id = ?", parentId).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetAll() ([]*models.Category, error) {
	var categories []*models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) Save(category *models.Category) error {
	if category.Id == 0 {
		return r.db.Create(category).Error
	} else if err := r.db.First(&category, category.Id).Error; err != nil {
		return err
	}
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint, hard bool) error {
	if hard {
		return r.db.Unscoped().Delete(&models.Category{}, id).Error
	}
	return r.db.Delete(&models.Category{}, id).Error
}

// Repository constructor returns the interface
func NewCategoryRepository(db *gorm.DB) CategoryRepositoryI {
	return &CategoryRepository{db: db}
}
