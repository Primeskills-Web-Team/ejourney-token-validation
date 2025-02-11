package auth

import (
    "github.com/golang-jwt/jwt"
)

type TokenClaims struct {
    Aud string `json:"aud"`
    Sub string `json:"sub"`
    jwt.StandardClaims
}

type TokenResponse struct {
    AccessToken  string `json:"access_token"`
    ExpiresIn   int    `json:"expires_in"`
    TokenType   string `json:"token_type"`
    RefreshToken string `json:"refresh_token"`
}

type User struct {
    ID       int    `json:"id"`
    FullName string `json:"full_name"`
    Email    string `json:"email"`
    // ... sesuaikan dengan response user yang ada
}