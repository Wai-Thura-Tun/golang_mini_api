package util

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(key string, userID uint64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

type RefreshToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func CreateRefreshToken() (*RefreshToken, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	tokenStr := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
	return &RefreshToken{
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}, nil
}
