// validator.go
package auth

import (
    "fmt"
    "strings"
    "github.com/golang-jwt/jwt"
)

type TokenValidator struct {
    secretKey string
}

func NewTokenValidator(secretKey string) *TokenValidator {
    return &TokenValidator{
        secretKey: secretKey,
    }
}

func (v *TokenValidator) ValidateToken(tokenString string) (*TokenClaims, error) {
    // Remove Bearer prefix if exists
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        // Validate signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(v.secretKey), nil
    })

    if err != nil {
        return nil, fmt.Errorf("failed to parse token: %w", err)
    }

    if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, ErrInvalidToken
}