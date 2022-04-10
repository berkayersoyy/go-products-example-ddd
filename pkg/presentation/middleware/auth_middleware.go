package middleware

import (
	"fmt"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeJWTMiddleware(a domain.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// const BEARER_SCHEMA = "Bearer"
		// authHeader := c.GetHeader("Authorization")
		// tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		token, err := a.ValidateToken(c.Request)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
