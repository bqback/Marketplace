package middleware

import (
	"context"
	"marketplace/internal/apperrors"
	"marketplace/internal/auth"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/utils"
	"net/http"
)

type AuthMiddleware struct {
	manager *auth.AuthManager
}

func NewAuthMiddleware(manager *auth.AuthManager) AuthMiddleware {
	return AuthMiddleware{manager: manager}
}

func (am AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, requestID, err := utils.GetLoggerAndID(r.Context())
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		funcName := "Auth"

		token := r.Header.Get("Authorization")
		if token == "" {
			logger.Error("Unauthorized access")
			apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
			return
		}
		logger.DebugFmt("Token found", requestID, funcName, nodeName)
		userId, err := am.manager.ValidateToken(token)
		if err != nil {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
			logger.Error("Invalid authorization")
			apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
			return
		}
		logger.DebugFmt("Token validated", requestID, funcName, nodeName)

		rCtx := context.WithValue(r.Context(), dto.UserIDKey, userId)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}

func (am AuthMiddleware) OptionalAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, requestID, err := utils.GetLoggerAndID(r.Context())
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		funcName := "Auth"

		rCtx := context.WithValue(r.Context(), dto.IsAuthorizedKey, true)

		token := r.Header.Get("Authorization")
		if token == "" {
			rCtx = context.WithValue(r.Context(), dto.IsAuthorizedKey, false)
		}
		logger.DebugFmt("Token found", requestID, funcName, nodeName)
		userId, err := am.manager.ValidateToken(token)
		if err != nil {
			rCtx = context.WithValue(r.Context(), dto.IsAuthorizedKey, false)
		}
		logger.DebugFmt("Token validated", requestID, funcName, nodeName)

		rCtx = context.WithValue(rCtx, dto.UserIDKey, userId)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
