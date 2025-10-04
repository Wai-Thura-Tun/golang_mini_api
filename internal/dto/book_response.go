package dto

type BookResponse struct {
	ID            uint64  `json:"id"`
	Name          string  `json:"name"`
	Overview      string  `json:"overview"`
	Type          string  `json:"type"`
	Cover         string  `json:"cover"`
	Author_Name   string  `json:"author_name"`
	Category_Name string  `json:"category_name"`
	Author_Bio    string  `json:"author_bio"`
	Rating        uint    `json:"rating"`
	Price         float64 `json:"price"`
	IsSpecial     bool    `json:"is_special"`
}
