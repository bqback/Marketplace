package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrNilContext = errors.New("context is nil")
)

var (
	ErrInvalidLoggingLevel = errors.New("invalid logging level")
	ErrLoggerMissing       = errors.New("logger missing from context")
)

var (
	ErrRequestIDMissing   = errors.New("request ID is missing from context")
	ErrFeedOptionsMissing = errors.New("sort options are missing from context")
)

var (
	ErrCouldNotParseURLParam = errors.New("failed to parse URL params")
	ErrUserIDMissing         = errors.New("user ID is missing from context")
)

var (
	ErrEnvNotFound         = errors.New("unable to load .env file")
	ErrDatabaseUserMissing = errors.New("database user is missing from env")
	ErrDatabasePWMissing   = errors.New("database password is missing from env")
	ErrDatabaseNameMissing = errors.New("database name is missing from env")
	ErrJWTSecretMissing    = errors.New("JWT secret is missing from env")
)

var (
	ErrCouldNotBuildQuery       = errors.New("failed to build SQL query")
	ErrCouldNotPrepareStatement = errors.New("failed to prepare query statement")
	ErrCouldNotBeginTransaction = errors.New("failed to start DB transaction")
	ErrCouldNotRollback         = errors.New("failed to roll back after a failed query")
	ErrCouldNotCommit           = errors.New("failed to commit DB transaction changes")
	ErrEmptyResult              = errors.New("no results for provided query")
)

var (
	ErrListingNotCreated  = errors.New("failed to insert listing into database")
	ErrListingNotSelected = errors.New("failed to select listing from database")
)

var (
	ErrCouldNotLinkListing = errors.New("failed to link listing to user")
)

var (
	ErrUserNotCreated  = errors.New("failed to insert user into database")
	ErrUserNotSelected = errors.New("failed to select user")
	ErrWrongPassword   = errors.New("wrong password")
)

var (
	ErrCouldNotParseClaims = errors.New("failed to parse JWT claims")
	ErrTokenExpired        = errors.New("token expired")
	ErrInvalidIssuedTime   = errors.New("invalid issue timestamp")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var BadRequestResponse = ErrorResponse{
	Code:    http.StatusBadRequest,
	Message: http.StatusText(http.StatusBadRequest),
}

var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: http.StatusText(http.StatusInternalServerError),
}

var UnauthorizedResponse = ErrorResponse{
	Code:    http.StatusUnauthorized,
	Message: http.StatusText(http.StatusUnauthorized),
}

func ReturnError(err ErrorResponse, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(err.Code)
	_, writeErr := w.Write([]byte(err.Message))
	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
