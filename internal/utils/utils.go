package utils

import (
	"context"
	"log"
	"marketplace/internal/apperrors"
	"marketplace/internal/logging"
	"marketplace/internal/pkg/dto"

	"golang.org/x/crypto/bcrypt"
)

func GetReqLogger(ctx context.Context) (logging.ILogger, error) {
	if ctx == nil {
		return nil, apperrors.ErrNilContext
	}
	if logger, ok := ctx.Value(dto.LoggerKey).(logging.ILogger); ok {
		return logger, nil
	}
	return nil, apperrors.ErrLoggerMissing
}

func GetReqID(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", apperrors.ErrNilContext
	}
	if reqID, ok := ctx.Value(dto.RequestIDKey).(string); ok {
		return reqID, nil
	}
	return "", apperrors.ErrRequestIDMissing
}

func GetUserID(ctx context.Context) (uint64, error) {
	if ctx == nil {
		return 0, apperrors.ErrNilContext
	}
	if id, ok := ctx.Value(dto.UserIDKey).(uint64); ok {
		return id, nil
	}
	return 0, apperrors.ErrUserIDMissing
}

func GetFeedOpts(ctx context.Context) (dto.FeedOptions, error) {
	if ctx == nil {
		return dto.FeedOptions{}, apperrors.ErrNilContext
	}
	if opts, ok := ctx.Value(dto.FeedOptionsKey).(dto.FeedOptions); ok {
		return opts, nil
	}
	return dto.FeedOptions{}, apperrors.ErrFeedOptionsMissing
}

func GetLoggerAndID(ctx context.Context) (logging.ILogger, string, error) {
	logger, err := GetReqLogger(ctx)
	if err != nil {
		log.Println(apperrors.ErrLoggerMissing)
		return nil, "", err
	}
	requestID, err := GetReqID(ctx)
	if err != nil {
		logger.Error(err.Error())
		return nil, "", err
	}
	return logger, requestID, nil
}

func HashFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePasswords(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
