package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Refresh token

type RefreshToken struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	IsRevoked bool      `json:"is_revoked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
