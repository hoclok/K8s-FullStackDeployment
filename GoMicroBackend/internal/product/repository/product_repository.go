package repository

import (
	"GoMicroBackend/internal/product/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Delete(id int64) error {
	return r.db.Delete(&model.Product{}, id).Error
}

func (r *ProductRepository) GetAll() ([]*model.Product, error) {
	var products []*model.Product
	err := r.db.Find(&products).Error
	return products, err
}
