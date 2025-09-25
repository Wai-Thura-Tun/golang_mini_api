package repos

import (
	"context"
	"database/sql"
)

type AuthorRepository interface {
	RegisterAuthor(ctx context.Context, name string, bio string) error
}

type AuthorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) *AuthorRepo {
	return &AuthorRepo{db: db}
}

func (repo *AuthorRepo) RegisterAuthor(ctx context.Context, name string, bio string) error {
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

	query := "insert into authors (name, biography) values (?, ?)"

	_, err = tx.ExecContext(ctx, query, name, bio)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
