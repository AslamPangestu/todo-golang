package lib

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTInteractor interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type initializeJWT struct {
}

func InitializeJWT() *initializeJWT {
	return &initializeJWT{}
}

var SECRET_KEY = []byte(os.Getenv("SECRET_JWT"))

func (s *initializeJWT) GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["aud"] = os.Getenv("FRONTEND_URL")
	claims["iss"] = os.Getenv("BACKEND_URL")
	claims["sub"] = userID
	claims["exp"] = time.Now().Unix() + 60*3600

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *initializeJWT) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("INVALID_TOKEN")
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
