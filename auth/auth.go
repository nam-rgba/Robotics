package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	Claims "github.com/nam-rgba/blv/models"
)

var jwtKey = []byte(os.Getenv("JWT"))

func GenerateToken(email string) (string, error) {
	claims := Claims.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		}, // 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (*Claims.Claims, error) {
	claims := &Claims.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &Claims.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims.Claims); ok && token.Valid {
		return claims, nil
	}
	return claims, nil
}
