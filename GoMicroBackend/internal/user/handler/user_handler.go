package handler

import (
	"GoMicroBackend/internal/user/service"
	"encoding/json"
	"net/http"
)

// @title User Service API
// @version 1.0
// @description User service with authentication endpoints
// @host localhost:8080
// @BasePath /api

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type RegisterRequest struct {
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"secret123"`
}

type LoginRequest struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"secret123"`
}

// @Summary Register a new user
// @Description Register a new user with username, email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register credentials"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {string} string "Invalid request or user already exists"
// @Router /register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary Login user
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} model.User
// @Failure 401 {string} string "Invalid credentials"
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// TODO: Generate JWT token here
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Logout user
// @Description Logout the current user
// @Tags auth
// @Produce json
// @Success 200 {string} string "Logged out successfully"
// @Router /logout [post]
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement token invalidation
	w.WriteHeader(http.StatusOK)
}
