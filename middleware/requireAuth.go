package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mayrista16/rest-api-postgres/database"
	"github.com/mayrista16/rest-api-postgres/models"
	"github.com/mayrista16/rest-api-postgres/services"
)

func RequireAuth(c *gin.Context) {
	// Get cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var request struct {
		AccessToken  string `json:"access_token" binding:"required"`
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	fmt.Println(request)
	// request validation should occur here
	userClaims := services.ParseAccessToken(request.AccessToken)
	refreshClaims := services.ParseRefreshToken(request.RefreshToken)
	fmt.Println(userClaims)
	// refresh token is expired
	if !refreshClaims.ExpiresAt.Time.After(time.Now()) {
		request.RefreshToken, err = services.NewRefreshToken(*refreshClaims)
		if err != nil {
			log.Fatal("error creating refresh token")
		}
	}

	// access token is expired
	if !userClaims.ExpiresAt.Time.After(time.Now()) && refreshClaims.ExpiresAt.Time.After(time.Now()) {
		request.AccessToken, err = services.NewAccessToken(*userClaims)
		if err != nil {
			log.Fatal("error creating access token")
		}
	}

	// Decode/validate
	secretKey := "4lly0uRth1n6S8eLoN6T0u5"
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find user with token sub
		var user models.User
		database.DB.First(&user, "id = ?", claims["sub"])

		if user.ID == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
