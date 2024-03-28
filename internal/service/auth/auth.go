package auth

import (
	"context"
	"marketplace/internal/auth"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/storage"
	"marketplace/internal/utils"
)

type AuthService struct {
	am *auth.AuthManager
	as storage.IAuthStorage
}

func NewAuthService(authStorage storage.IAuthStorage, manager *auth.AuthManager) *AuthService {
	return &AuthService{
		am: manager,
		as: authStorage,
	}
}

func (s *AuthService) Authorize(ctx context.Context, info dto.LoginInfo) (*dto.JWT, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "Authenticate"

	user, err := s.as.Auth(ctx, info)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Login info correct", requestID, funcName, "service")
	tokenString, err := s.am.GenerateToken(user)
	if err != nil {
		logger.DebugFmt("Token not generated "+err.Error(), requestID, funcName, "service")
		return nil, err
	}
	return &dto.JWT{Token: tokenString}, nil
}

func (s *AuthService) Register(ctx context.Context, info dto.SignupInfo) (*dto.NewUserInfo, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "Register"

	hash, err := utils.HashFromPassword(info.Password)
	if err != nil {
		return nil, err
	}

	info.Password = hash
	user, err := s.as.Register(ctx, info)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("User registered", requestID, funcName, "service")

	tokenString, err := s.am.GenerateToken(user)
	if err != nil {
		logger.DebugFmt("Token not generated "+err.Error(), requestID, funcName, "service")
		return nil, err
	}

	fullUser := dto.NewUserInfo{
		ID:    user.ID,
		Login: user.Login,
		Token: tokenString,
	}
	return &fullUser, nil
}
