package service

import (
	"context"
	"time"

	"GoMicroBackend/internal/product/model"
	"GoMicroBackend/internal/product/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	return s.repo.Delete(id)
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]*model.Product, error) {
	return s.repo.GetAll()
}
