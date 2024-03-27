package auth

import (
	"marketplace/internal/apperrors"
	"marketplace/internal/config"
	"marketplace/internal/pkg/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthManager struct {
	secret   []byte
	lifetime time.Duration
}

type jwtClaim struct {
	ID   uint64
	Name string
	jwt.RegisteredClaims
}

func NewManager(config *config.JWTConfig) *AuthManager {
	return &AuthManager{secret: []byte(config.Secret), lifetime: config.Lifetime}
}

func (am *AuthManager) GenerateToken(user *entities.User) (string, error) {
	expiresAt := time.Now().Add(am.lifetime)
	claims := &jwtClaim{
		ID:   user.ID,
		Name: user.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(am.secret)
}

func (am *AuthManager) ValidateToken(token string) (uint64, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(am.secret), nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(*jwtClaim)
	if !ok {
		return 0, apperrors.ErrCouldNotParseClaims
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		return 0, apperrors.ErrTokenExpired
	}
	if claims.IssuedAt.After(time.Now().Local()) {
		return 0, apperrors.ErrInvalidIssuedTime
	}

	return claims.ID, nil
}
