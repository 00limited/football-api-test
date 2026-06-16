package middleware

import (
	"net/http"
	"strings"

	"github.com/00limited/football-api/internal/config"
	resp "github.com/00limited/football-api/internal/dto/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWT(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp.APIResponse{Status: "error", Message: "unauthorized", Errors: []string{"missing bearer token"}})
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp.APIResponse{Status: "error", Message: "unauthorized", Errors: []string{"invalid token"}})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp.APIResponse{Status: "error", Message: "unauthorized", Errors: []string{"invalid token claims"}})
			return
		}
		c.Set("admin_id", claims["sub"])
		c.Set("username", claims["username"])
		c.Next()
	}
}
