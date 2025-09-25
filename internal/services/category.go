package services

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/repos"
	"github.com/gofiber/fiber/v2"
)

type CategoryService struct {
	repo *repos.CategoryRepo
}

func NewCategoryService(repo *repos.CategoryRepo) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (service *CategoryService) RegisterCategory(ctx *fiber.Ctx, req *dto.CategoryRegisterRequest) *dto.Response {
	resp := &dto.Response{
		Obj: make(map[string]string),
	}
	err := service.repo.RegisterCategory(ctx.Context(), req)
	if err != nil {
		resp.Code = fiber.StatusInternalServerError
		resp.Obj = map[string]string{
			"error": "Something went wrong",
		}
		return resp
	}
	resp.Code = fiber.StatusCreated
	resp.Obj = map[string]string{
		"message": "Category has been created.",
	}
	return resp
}
