package repos

import (
	"context"
	"database/sql"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
)

type CategoryRepository interface {
	RegisterCategory(ctx context.Context, req *dto.CategoryRegisterRequest) error
}

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (repo *CategoryRepo) RegisterCategory(ctx context.Context, req *dto.CategoryRegisterRequest) error {
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

	query := "insert into category (name) values (?)"
	_, err = tx.ExecContext(ctx, query, req.Name)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
