package middleware

import (
    "net/http"
    "fmt"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)



func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		//var jwtSecretKey = []byte("JWT_SECRET")
		jwtSecretKey := []byte(os.Getenv("JWT_SECRET"))

		fmt.Println("Secret in Middleware:", string(jwtSecretKey))

        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        /*token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecretKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }*/
        // ✅ Handle “Bearer <token>” format
        parts := strings.Split(authHeader, " ")
        var tokenString string
        if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
            tokenString = parts[1]
        } else {
            tokenString = authHeader
        }

        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecretKey, nil
        })

        if err != nil || !token.Valid {
            fmt.Println("Token parse error:", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Optional: print claims for debugging
        fmt.Println("Token claims:", claims)


        c.Next()
    }
}
