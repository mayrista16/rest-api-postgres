package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": claims.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte("4lly0uRth1n6S8eLoN6T0u5")
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		panic(err)
	}

	return tokenString, err

}

func NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte("4lly0uRth1n6S8eLoN6T0u5"))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("4lly0uRth1n6S8eLoN6T0u5"), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("4lly0uRth1n6S8eLoN6T0u5"), nil
	})

	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
}
