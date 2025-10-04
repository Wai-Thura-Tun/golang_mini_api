package dto

type CategoryRegisterRequest struct {
	Name string `json:"name" validate:"required"`
}
