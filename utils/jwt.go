package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	godotenv.Load()

	secretKey = []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0 {
		secretKey = []byte("secret")
	}
}
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username":   username,
		"expired_at": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyJWT(token string) (jwt.MapClaims, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
		if expired, ok := claims["expired_at"].(float64); ok {
			if time.Now().Unix() > int64(expired) {
				return nil, errors.New("token expired")
			}
		}

		return claims, nil
	}

	return nil, errors.New("invalid token")
}
