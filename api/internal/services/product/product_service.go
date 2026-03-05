package product

import (
	dto "commerce/api/internal/dto/product"
	repo "commerce/internal/shared/repositories/product"
)

type ProductServiceI interface {
	GetById(id uint) (*dto.Product, error)
	GetAll() ([]*dto.Product, error)
	GetAllByCategory(categoryId uint) ([]*dto.Product, error)
	GetAllByOrder(orderId uint) ([]*dto.Product, error)
	Save(product *dto.Product) error
	Delete(id uint, hard bool) error
}

type ProductService struct {
	repo repo.ProductRepositoryI
}

func NewProductService(repo repo.ProductRepositoryI) ProductServiceI {
	return &ProductService{repo: repo}
}

// Delete implements [ProductServiceI].
func (p *ProductService) Delete(id uint, hard bool) error {
	return p.repo.Delete(id, hard)
}

// GetAll implements [ProductServiceI].
func (p *ProductService) GetAll() ([]*dto.Product, error) {
	panic("unimplemented")
}

// GetAllByCategory implements [ProductServiceI].
func (p *ProductService) GetAllByCategory(categoryId uint) ([]*dto.Product, error) {
	panic("unimplemented")
}

// GetAllByOrder implements [ProductServiceI].
func (p *ProductService) GetAllByOrder(orderId uint) ([]*dto.Product, error) {
	panic("unimplemented")
}

// GetById implements [ProductServiceI].
func (p *ProductService) GetById(id uint) (*dto.Product, error) {
	panic("unimplemented")
}

// Save implements [ProductServiceI].
func (p *ProductService) Save(product *dto.Product) error {
	model := dto.ToModel(product)
	return p.repo.Save(model)
}
