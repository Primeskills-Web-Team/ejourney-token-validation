package auth

import "errors"

var (
    ErrInvalidToken        = errors.New("invalid token")
    ErrTokenExpired        = errors.New("token has expired")
    ErrMissingAuthHeader   = errors.New("missing authorization header")
    ErrInvalidAuthHeader   = errors.New("invalid authorization header format")
)

type AuthError struct {
    Code        string `json:"code"`
    Message     string `json:"message"`
    IsSuccess   bool   `json:"is_success"`
    ErrorDetail string `json:"error_detail"`
    Resource    string `json:"resource"`
}

func NewAuthError(code, message, detail, resource string) AuthError {
    return AuthError{
        Code:        code,
        Message:     message,
        IsSuccess:   false,
        ErrorDetail: detail,
        Resource:    resource,
    }
}