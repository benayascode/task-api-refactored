package infrastructure

import (
	"time"

	domain "task-manager/Domain"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("BlackBox")

type JWTService interface {
	GenerateToken(user domain.User) (string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}

type jwtServiceImpl struct{}

func NewJWTService() JWTService {
	return &jwtServiceImpl{}
}

func (j *jwtServiceImpl) GenerateToken(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.UserName,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (j *jwtServiceImpl) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
