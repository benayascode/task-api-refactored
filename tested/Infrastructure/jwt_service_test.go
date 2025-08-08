package infrastructure

import (
	"testing"
	"time"

	domain "task-manager/Domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJWTService_GenerateAndValidate(t *testing.T) {
	service := NewJWTService()
	user := domain.User{UserName: "bob", Role: "Admin"}

	token, err := service.GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := service.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "bob", claims["username"])
	assert.Equal(t, "Admin", claims["role"])
}

func TestJWTService_ValidateToken_Invalid(t *testing.T) {
	service := NewJWTService()
	_, err := service.ValidateToken("invalid.token.value")
	assert.Error(t, err)
	assert.ErrorIs(t, err, jwt.ErrTokenMalformed)
}

func TestJWTService_ValidateToken_Expired(t *testing.T) {
	service := NewJWTService()
	// Create a token already expired
	claims := jwt.MapClaims{
		"username": "bob",
		"role":     "user",
		"exp":      time.Now().Add(-time.Hour).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := jwtToken.SignedString([]byte("BlackBox"))
	assert.NoError(t, err)

	_, err = service.ValidateToken(signed)
	assert.Error(t, err)
	assert.ErrorIs(t, err, jwt.ErrTokenExpired)
}
