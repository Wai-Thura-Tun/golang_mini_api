package services

import (
	"context"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/models"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/repos"
)

type AuthService struct {
	repo *repos.AuthRepo
}

func NewAuthService(repo *repos.AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (service *AuthService) CreateUser(ctx context.Context, req dto.RegisterRequest) error {
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	user.HashPassword()
	return service.repo.CreateUser(ctx, user)
}
