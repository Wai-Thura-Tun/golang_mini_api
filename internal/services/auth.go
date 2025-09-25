package services

import (
	"time"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/models"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/repos"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/util"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repos.AuthRepo
}

func NewAuthService(repo *repos.AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (service *AuthService) CreateUser(ctx *fiber.Ctx, req *dto.RegisterRequest) *dto.Response {
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	user.HashPassword()
	isExist, err := service.repo.CheckExistEmail(ctx.Context(), req.Email)

	respObj := &dto.Response{
		Obj: make(map[string]string),
	}
	if err != nil {
		respObj.Code = fiber.StatusInternalServerError
		respObj.Obj = map[string]string{
			"error": "Unable to check email",
		}
	}
	if isExist {
		respObj.Code = fiber.StatusConflict
		respObj.Obj = map[string]string{
			"error": "Your email is already registered. Please use different email",
		}
	} else {
		err = service.repo.CreateUser(ctx.Context(), user)
		if err != nil {
			respObj.Code = fiber.StatusInternalServerError
			respObj.Obj = map[string]string{
				"error": "Unable to register your email",
			}
		} else {
			respObj.Code = fiber.StatusCreated
			respObj.Obj = map[string]string{
				"message": "Your email is registered successfully",
			}
		}
	}
	return respObj
}

func (service *AuthService) Login(ctx *fiber.Ctx, req *dto.LoginRequest, secret string) *dto.Response {
	respObj := &dto.Response{
		Obj: make(map[string]string),
	}
	userID, pass, err := service.repo.GetUserCredentials(ctx.Context(), req.Email)
	if err != nil {
		respObj.Code = fiber.StatusNotFound
		respObj.Obj = map[string]string{
			"error": "Your email may not be registered yet",
		}
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(req.Password))
		if err != nil {
			respObj.Code = fiber.StatusUnauthorized
			respObj.Obj = map[string]string{
				"error": "Your password is wrong",
			}
		} else {
			err = service.repo.RevokeOldRefreshToken(ctx.Context(), userID)
			if err != nil {
				respObj.Code = fiber.StatusInternalServerError
				respObj.Obj = map[string]string{
					"error": "Something went wrong. Try again later",
				}
				return respObj
			}
			refreshTokenStr, err := util.CreateRefreshToken()
			if err != nil {
				respObj.Code = fiber.StatusInternalServerError
				respObj.Obj = map[string]string{
					"error": "Something went wrong",
				}
				return respObj
			}
			err = service.repo.StoreRefreshToken(ctx.Context(), refreshTokenStr, userID)
			if err != nil {
				respObj.Code = fiber.StatusInternalServerError
				respObj.Obj = map[string]string{
					"error": "Unable to login",
				}
			} else {
				tokenStr, err := util.CreateToken(secret, userID)
				if err != nil {
					respObj.Code = fiber.StatusInternalServerError
					respObj.Obj = map[string]string{
						"error": "Unable to login",
					}
				} else {
					respObj.Code = fiber.StatusOK
					respObj.Obj = map[string]string{
						"message":       "Login successful",
						"access_token":  tokenStr,
						"refresh_token": refreshTokenStr.Token,
					}
				}
			}
		}
	}
	return respObj
}

func (service *AuthService) RefreshAccessToken(ctx *fiber.Ctx, req *dto.RefreshRequest, secret string) *dto.Response {
	respObj := &dto.Response{
		Obj: make(map[string]string),
	}
	userId, expiredAt, err := service.repo.FetchRefreshTokenInfo(ctx.Context(), req.RefreshToken)
	if err != nil {
		respObj.Code = fiber.StatusUnauthorized
		respObj.Obj = map[string]string{
			"error": "Token is invalid",
		}
	} else {
		if expiredAt.Before(time.Now()) {
			respObj.Code = fiber.StatusUnauthorized
			respObj.Obj = map[string]string{
				"error": "Authentication is expired",
			}
		} else {
			refreshToken, err := util.CreateRefreshToken()
			if err != nil {
				respObj.Code = fiber.StatusInternalServerError
				respObj.Obj = map[string]string{
					"error": "Something went wrong",
				}
				return respObj
			}

			err = service.repo.RotateRefreshToken(ctx.Context(), req.RefreshToken, refreshToken, userId)
			if err != nil {
				respObj.Code = fiber.StatusInternalServerError
				respObj.Obj = map[string]string{
					"error": "Something went wrong",
				}
			} else {
				accessToken, err := util.CreateToken(secret, userId)
				if err != nil {
					respObj.Code = fiber.StatusInternalServerError
					respObj.Obj = map[string]string{
						"error": "Something went wrong",
					}
					return respObj
				}
				respObj.Code = fiber.StatusOK
				respObj.Obj = map[string]string{
					"message":       "Access Token refresh successfully",
					"access_token":  accessToken,
					"refresh_token": refreshToken.Token,
				}
			}
		}
	}
	return respObj
}

func (service *AuthService) Logout(ctx *fiber.Ctx, userID uint64) *dto.Response {
	respObj := &dto.Response{
		Obj: map[string]string{},
	}
	err := service.repo.RevokeOldRefreshToken(ctx.Context(), userID)
	if err != nil {
		respObj.Code = fiber.StatusInternalServerError
		respObj.Obj = map[string]string{
			"error": "Something went wrong",
		}
	}
	respObj.Code = fiber.StatusOK
	respObj.Obj = map[string]string{
		"message": "You've been logout successfully",
	}
	return respObj
}
