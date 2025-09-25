package dto

type AuthorRegisterRequest struct {
	Name      string `json:"name" validate:"required"`
	Biography string `json:"biography" validate:"required"`
}
