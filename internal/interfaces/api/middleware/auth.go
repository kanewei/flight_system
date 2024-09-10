package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthRole string

const (
	Staff AuthRole = "staff"
	User  AuthRole = "user"
)

var signingKey = []byte("your_secret_key_here")

func AuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const apiKey = "your_api_key_here"
		// Check for API key in request headers
		if key := ctx.GetHeader("API-Key"); key != apiKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		ctx.Next()
	}
}

func StaffAuthMiddlerware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		splitToken := strings.Split(tokenString, " ")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			ctx.Abort()
			return
		}

		tokenString = splitToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return signingKey, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if role, ok := claims["role"].(string); ok && role == string(Staff) {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized for this resource"})
		ctx.Abort()
	}
}

func UserAuthMiddlerware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		splitToken := strings.Split(tokenString, " ")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			ctx.Abort()
			return
		}

		tokenString = splitToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return signingKey, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if role, ok := claims["role"].(string); ok && role == string(User) {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized for this resource"})
		ctx.Abort()
	}
}
