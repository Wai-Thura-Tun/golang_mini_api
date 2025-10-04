package repos

import (
	"context"
	"database/sql"
	"log"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
)

type BookRepository interface {
	RegisterBook(ctx context.Context, req *dto.BookRequest) error
	GetBooks(ctx context.Context) error
}

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (repo *BookRepo) RegisterBook(ctx context.Context, req *dto.BookRequest) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	query := "insert into books (name, overview, type, cover, author_id, category_id, rating, price, isSpecial) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, req.Name, req.Overview, req.Type, req.Cover, req.AuthorID, req.CategoryID, req.Rating, req.Price, req.IsSpecial)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (repo *BookRepo) GetBooks(ctx context.Context) ([]*dto.BookResponse, error) {
	query := `select b.id, b.name, b.overview, b.type, b.cover, b.rating, b.price, b.isSpecial AS is_special, a.name AS author_name, a.biography AS author_bio, c.name AS category_name
			  from books b
			  join authors a on b.author_id = a.id
			  join category c on b.category_id = c.id;
	         `

	books := make([]*dto.BookResponse, 0)

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(dto.BookResponse)

		if err := rows.Scan(&book.ID, &book.Name, &book.Overview, &book.Type, &book.Cover, &book.Rating, &book.Price, &book.IsSpecial, &book.Author_Name, &book.Author_Bio, &book.Category_Name); err != nil {
			log.Println(err.Error())
			return books, err
		}

		books = append(books, book)

	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return books, err
	}

	return books, nil
}
