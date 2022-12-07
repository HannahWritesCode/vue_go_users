package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"vue-go-users.com/models"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the token
		tokenString, errorCode := getToken(c)
		if errorCode == http.StatusUnauthorized {
			c.AbortWithStatus(errorCode)
			return
		}

		// Check that token is not empty
		if len(tokenString) == 0 {
			c.AbortWithStatus(401)
			return
		}

		// Create struct to store claims in
		type Claims struct {
			Token string `json:"token"`
			ID    uint   `json:"id"`
			jwt.StandardClaims
		}
		claims := &Claims{}

		// Validate JWT
		token, parseErr := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APP_KEY")), nil
		})
		if parseErr != nil {
			if parseErr == jwt.ErrSignatureInvalid {
				c.AbortWithStatus(401)
				return
			}
			c.AbortWithStatus(400)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			c.AbortWithStatus(401)
			return
		}

		// Find user by their token
		user, getByTokenErr := models.User{}.GetByAuthToken(claims.ID, claims.Token)
		if getByTokenErr != nil {
			c.AbortWithStatus(422)
			return
		}

		c.Set("user", user)

		c.Next()
	}

}

// getToken gets the Bearer JWT
func getToken(c *gin.Context) (string, int) {
	// Get the logged in token
	reqAuth := c.Request.Header.Get("Authorization")

	authSplit := strings.Split(reqAuth, "Bearer")

	if len(authSplit) != 2 {
		return "", http.StatusUnauthorized
	}

	return strings.TrimSpace(authSplit[1]), 0
}
