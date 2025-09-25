package repos

import (
	"context"
	"database/sql"
	"time"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/models"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/util"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	CheckExistEmail(ctx context.Context, email string) (bool, error)
}

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (repo *AuthRepo) CreateUser(ctx context.Context, user *models.User) error {
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

	query := "insert into users (name, email, password) values (?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, user.Name, user.Email, user.Password)

	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (repo *AuthRepo) CheckExistEmail(ctx context.Context, email string) (bool, error) {
	var exist bool
	query := "select 1 from users where email = ? limit 1"
	err := repo.db.QueryRowContext(ctx, query, email).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (repo *AuthRepo) GetUserCredentials(ctx context.Context, email string) (uint64, string, error) {
	query := "select id, password from users where email = ? limit 1"
	var pass string
	var id uint64
	row := repo.db.QueryRowContext(ctx, query, email)
	if err := row.Scan(&id, &pass); err != nil {
		return 0, "", err
	}
	return id, pass, nil
}

func (repo *AuthRepo) StoreRefreshToken(ctx context.Context, refreshToken *util.RefreshToken, userID uint64) error {
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

	query := "insert into refresh_tokens (user_id, token, expires_at) values (?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, userID, refreshToken.Token, refreshToken.ExpiresAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (repo *AuthRepo) RevokeOldRefreshToken(ctx context.Context, userId uint64) error {
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

	query := "update refresh_tokens set revoked=1 where user_id=? and revoked=0"
	_, err = tx.ExecContext(ctx, query, userId)

	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (repo *AuthRepo) FetchRefreshTokenInfo(ctx context.Context, refreshToken string) (uint64, time.Time, error) {
	query := "select user_id, expires_at from refresh_tokens where token=? and revoked=0 limit 1"
	var userId uint64
	var expiresAt time.Time
	row := repo.db.QueryRowContext(ctx, query, refreshToken)
	if err := row.Scan(&userId, &expiresAt); err != nil {
		return 0, time.Now(), err
	}
	return userId, expiresAt, nil
}

func (repo *AuthRepo) RotateRefreshToken(ctx context.Context, oldToken string, refreshToken *util.RefreshToken, userID uint64) error {
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

	_, err = tx.ExecContext(
		ctx,
		"update refresh_tokens set revoked=1 where token=?",
		oldToken,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"insert into refresh_tokens (user_id, token, expires_at) values (?, ?, ?)",
		userID,
		refreshToken.Token,
		refreshToken.ExpiresAt,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
