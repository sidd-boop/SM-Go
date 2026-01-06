package middleware

import (
	
	"net/http"
	"strings"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		//get auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Authorization header missing"})
			c.Abort()
			return
		}
		//expected format: "Bearer <token>"
		parts:=strings.Split(authHeader," ")
		if len(parts)!=2 || parts[0]!="Bearer"{
			c.JSON (http.StatusUnauthorized, gin. H{"error":"Invalid Authorization header format"})
			c.Abort()
			return
		}
		//parse and verify token
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid or expired token"})
			c.Abort()
			return
		}

		//Extract claims
		claims, ok:=token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token claims"})
			c.Abort()
			return
		}

		userIDFloat, ok:=claims["user_id"].(float64)
		if !ok{
			c.JSON (http.StatusUnauthorized, gin.H{"error":"Invalid token claims"})
			c.Abort()
			return
		}

		userID:=uint(userIDFloat)


		//attach userID to context
		c.Set("userID", userID)

		c.Next()

	}
}