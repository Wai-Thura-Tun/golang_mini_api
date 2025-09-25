package services

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/repos"
	"github.com/gofiber/fiber/v2"
)

type AuthorService struct {
	repo *repos.AuthorRepo
}

func NewAuthorService(repo *repos.AuthorRepo) *AuthorService {
	return &AuthorService{repo: repo}
}

func (service *AuthorService) RegisterAuthor(ctx *fiber.Ctx, req *dto.AuthorRegisterRequest) *dto.Response {
	resp := &dto.Response{
		Obj: make(map[string]string),
	}
	err := service.repo.RegisterAuthor(ctx.Context(), req.Name, req.Biography)
	if err != nil {
		resp.Code = fiber.StatusInternalServerError
		resp.Obj = map[string]string{
			"error": "Something went wrong",
		}
		return resp
	}
	resp.Code = fiber.StatusCreated
	resp.Obj = map[string]string{
		"message": "Author has been created",
	}
	return resp
}
