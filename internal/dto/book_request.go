package dto

type BookRequest struct {
	Name       string  `json:"name" validate:"required"`
	Overview   string  `json:"overview" validate:"required"`
	Type       string  `json:"type" validate:"required"`
	Cover      string  `json:"cover" validate:"required"`
	AuthorID   uint64  `json:"author_id" validate:"required"`
	CategoryID uint64  `json:"category_id" validate:"required"`
	Rating     uint16  `json:"rating" validate:"required"`
	Price      float32 `json:"price" validate:"required"`
	IsSpecial  bool    `json:"is_special" validate:"required"`
}
