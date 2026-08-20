package middleware

/*import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "os"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            if claims["role"] != "admin" {
                c.JSON(http.StatusForbidden, gin.H{"error": "Admins only"})
                c.Abort()
                return
            }
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}*/
