package services

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/repos"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/util"
	"github.com/gofiber/fiber/v2"
)

type BookService struct {
	repo *repos.BookRepo
}

func NewBookService(repo *repos.BookRepo) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (service *BookService) RegisterBook(c *fiber.Ctx, req *dto.BookRequest) *dto.Response {
	resp := &dto.Response{
		Obj: make(map[string]interface{}),
	}
	err := service.repo.RegisterBook(c.Context(), req)
	if err != nil {
		resp.Code = fiber.StatusInternalServerError
		resp.Obj = map[string]interface{}{
			"error": "Something went wrong",
		}
		return resp
	}
	resp.Code = fiber.StatusOK
	resp.Obj = map[string]interface{}{
		"message": "Book has been added successfully.",
	}
	return resp
}

func (service *BookService) GetBooks(c *fiber.Ctx) *dto.Response {
	resp := &dto.Response{
		Obj: make(map[string]interface{}),
	}
	books, err := service.repo.GetBooks(c.Context())
	if err != nil {
		resp.Code = fiber.StatusInternalServerError
		resp.Obj = map[string]interface{}{
			"error": "Something went wrong",
		}
		return resp
	}
	resp.Code = fiber.StatusOK
	resp.Obj = map[string]interface{}{
		"data": map[string]interface{}{
			"special_books": util.Filter(books, func(item *dto.BookResponse) bool {
				return item.IsSpecial
			}),
			"normal_books": util.Filter(books, func(item *dto.BookResponse) bool {
				return !item.IsSpecial
			}),
		},
	}
	return resp
}
