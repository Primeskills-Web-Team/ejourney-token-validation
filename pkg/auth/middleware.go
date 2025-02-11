package auth

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware(validator *TokenValidator, serviceName string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, NewAuthError(
                "401",
                "Unauthorized",
                "Missing authorization header",
                serviceName,
            ))
            c.Abort()
            return
        }

        claims, err := validator.ValidateToken(authHeader)
        if err != nil {
            c.JSON(http.StatusUnauthorized, NewAuthError(
                "401",
                "Invalid token",
                err.Error(),
                serviceName,
            ))
            c.Abort()
            return
        }

        // Set user data to context
        c.Set("user_id", claims.Sub)
        c.Set("aud", claims.Aud)
        
        c.Next()
    }
}