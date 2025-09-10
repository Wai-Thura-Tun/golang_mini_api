package repos

import (
	"context"
	"database/sql"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/models"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (repo *AuthRepo) CreateUser(ctx context.Context, user *models.User) error {
	query := "insert into users (name, email, password) values ($1, $2, $3)"
	_, err := repo.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}
	return nil
}
