package Infrastructure

import (
	"task_manager/Domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretWord = []byte("BlackBox")

type SimpleJWTService struct{}

func NewSimpleJWTService() Domain.JWTService {
	return &SimpleJWTService{}
}

func (s *SimpleJWTService) GenerateToken(user Domain.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.UserName,
		"role":     user.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretWord)
}
