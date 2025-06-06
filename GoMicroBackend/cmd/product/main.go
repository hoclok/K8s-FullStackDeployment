package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"GoMicroBackend/api/proto/product"
	_ "GoMicroBackend/docs"
	"GoMicroBackend/internal/product/handler"
	"GoMicroBackend/internal/product/middleware"
	"GoMicroBackend/internal/product/model"
	"GoMicroBackend/internal/product/repository"
	"GoMicroBackend/internal/product/service"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Product Service API
// @version 1.0
// @description Product Service API with gRPC and HTTP endpoints
// @host localhost:8081
// @BasePath /api/v1

func main() {
	// Database connection
	dsn := "host=postgres user=postgres password=postgres dbname=productgodb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository, service and handler
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Create a listener on TCP port 50051 for gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server with authentication middleware
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	)
	product.RegisterProductServiceServer(grpcServer, productHandler)

	// Start gRPC server
	go func() {
		fmt.Println("gRPC server starting on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("Failed to serve gRPC:", err)
		}
	}()

	// Create HTTP router
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Product endpoints
	api.HandleFunc("/products", createProduct(productService)).Methods("POST")
	api.HandleFunc("/products", getAllProducts(productService)).Methods("GET")
	api.HandleFunc("/products/{id}", deleteProduct(productService)).Methods("DELETE")
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Start HTTP server
	fmt.Println("HTTP server starting on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("Failed to serve HTTP:", err)
	}
}

// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product details"
// @Success 200 {object} model.Product
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [post]
func createProduct(s *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product model.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdProduct, err := s.CreateProduct(r.Context(), &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdProduct)
	}
}

// @Summary Get all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func getAllProducts(s *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := s.GetAllProducts(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [delete]
func deleteProduct(s *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		err = s.DeleteProduct(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
