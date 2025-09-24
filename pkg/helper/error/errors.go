package errorhelper

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/pkg/errors"
)

// AppError represents an application error
type AppError struct {
	Code    string
	Message string
	Status  int
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code, message string, status int, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
		Err:     err,
	}
}

// Common error codes
const (
	ErrNotFound            = "NOT_FOUND"
	ErrBadRequest          = "BAD_REQUEST"
	ErrUnauthorized        = "UNAUTHORIZED"
	ErrForbidden           = "FORBIDDEN"
	ErrInternalServer      = "INTERNAL_SERVER_ERROR"
	ErrValidation          = "VALIDATION_ERROR"
	ErrDuplicateEntry      = "DUPLICATE_ENTRY"
	ErrDatabaseOperation   = "DATABASE_ERROR"
	ErrMissingDBConnection = "MISSING_DB_CONNECTION"
	ErrMissingID           = "MISSING_ID"
	ErrMissingUpdateMap    = "MISSING_UPDATE_MAP"
)

// NotFound creates a new not found error
func NotFound(message string, err error) *AppError {
	return NewAppError(ErrNotFound, message, http.StatusNotFound, err)
}

// BadRequest creates a new bad request error
func BadRequest(message string, err error) *AppError {
	return NewAppError(ErrBadRequest, message, http.StatusBadRequest, err)
}

// Unauthorized creates a new unauthorized error
func Unauthorized(message string, err error) *AppError {
	return NewAppError(ErrUnauthorized, message, http.StatusUnauthorized, err)
}

// Forbidden creates a new forbidden error
func Forbidden(message string, err error) *AppError {
	return NewAppError(ErrForbidden, message, http.StatusForbidden, err)
}

// InternalServer creates a new internal server error
func InternalServer(message string, err error) *AppError {
	return NewAppError(ErrInternalServer, message, http.StatusInternalServerError, err)
}

// MissingDBConnection creates a new missing DB connection error
func MissingDBConnection(message string, err error) *AppError {
	return NewAppError(ErrMissingDBConnection, message, http.StatusInternalServerError, err)
}

// MissingChapterID creates a new missing chapter ID error
func MissingID(message string, err error) *AppError {
	return NewAppError(ErrMissingID, message, http.StatusInternalServerError, err)
}

// MissingUpdateMap creates a new missing update map or translations error
func MissingUpdateMap(message string, err error) *AppError {
	return NewAppError(ErrMissingUpdateMap, message, http.StatusInternalServerError, err)
}

// Validation creates a new validation error
func Validation(message string, err error) *AppError {
	return NewAppError(ErrValidation, message, http.StatusUnprocessableEntity, err)
}

// Wrap wraps an error with a message
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func ComposeStacktrace(err error) *string {
	if err == nil {
		return nil
	}

	out := fmt.Sprintf("error: %+v\nstacktrace: %s", err, string(debug.Stack()))

	return &out
}
