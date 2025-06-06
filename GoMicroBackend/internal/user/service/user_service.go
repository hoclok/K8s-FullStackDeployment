package service

import (
	"context"
	"errors"
	"time"

	"GoMicroBackend/api/proto/product"
	"GoMicroBackend/internal/user/model"
	"GoMicroBackend/internal/user/repository"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	repo          *repository.UserRepository
	productClient product.ProductServiceClient
}

func NewUserService(repo *repository.UserRepository) *UserService {
	// Connect to product service
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	productClient := product.NewProductServiceClient(conn)

	return &UserService{
		repo:          repo,
		productClient: productClient,
	}
}

func (s *UserService) Register(username, email, password string) error {
	// Check if username exists
	if _, err := s.repo.FindByUsername(username); err == nil {
		return errors.New("username already exists")
	}

	// Check if email exists
	if _, err := s.repo.FindByEmail(email); err == nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(user)
}

func (s *UserService) Login(username, password string) (*model.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	return s.repo.Update(user)
}
