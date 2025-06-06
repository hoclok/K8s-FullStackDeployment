package main

import (
	_ "GoMicroBackend/docs" // swagger docs
	"GoMicroBackend/internal/user/handler"
	"GoMicroBackend/internal/user/model"
	"GoMicroBackend/internal/user/repository"
	"GoMicroBackend/internal/user/service"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title User Service API
// @version 1.0
// @description User service with authentication endpoints
// @host localhost:8080
// @BasePath /api
func main() {
	// Database connection
	dsn := "host=postgres user=postgres password=postgres dbname=usergodb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	http.HandleFunc("/api/register", userHandler.Register)
	http.HandleFunc("/api/login", userHandler.Login)
	http.HandleFunc("/api/logout", userHandler.Logout)
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Swagger documentation
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Start server
	fmt.Println("Server starting on :8080")
	fmt.Println("Swagger UI available at: http://localhost:8080/swagger/index.html")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
