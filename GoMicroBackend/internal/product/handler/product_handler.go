package handler

import (
	"context"
	"log"

	productpb "GoMicroBackend/api/proto/product"
	"GoMicroBackend/internal/product/model"
	"GoMicroBackend/internal/product/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductHandler struct {
	productpb.UnimplementedProductServiceServer
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.ProductResponse, error) {
	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	createdProduct, err := h.productService.CreateProduct(ctx, product)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return nil, status.Error(codes.Internal, "failed to create product")
	}

	return &productpb.ProductResponse{
		Product: &productpb.Product{
			Id:          createdProduct.ID,
			Name:        createdProduct.Name,
			Description: createdProduct.Description,
			Price:       createdProduct.Price,
			Stock:       createdProduct.Stock,
			CreatedAt:   createdProduct.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   createdProduct.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	err := h.productService.DeleteProduct(ctx, req.Id)
	if err != nil {
		log.Printf("Error deleting product: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete product")
	}

	return &productpb.DeleteProductResponse{
		Success: true,
	}, nil
}

func (h *ProductHandler) GetAllProducts(ctx context.Context, req *productpb.GetAllProductsRequest) (*productpb.GetAllProductsResponse, error) {
	products, err := h.productService.GetAllProducts(ctx)
	if err != nil {
		log.Printf("Error getting all products: %v", err)
		return nil, status.Error(codes.Internal, "failed to get products")
	}

	var productResponses []*productpb.Product
	for _, p := range products {
		productResponses = append(productResponses, &productpb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &productpb.GetAllProductsResponse{
		Products: productResponses,
	}, nil
}
