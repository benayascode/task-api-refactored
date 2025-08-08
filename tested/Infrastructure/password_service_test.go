package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordService_HashAndCompare(t *testing.T) {
	service := NewPasswordService()

	hash, err := service.HashPassword("kebadMister")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, "kebadMister", hash)

	// correct password
	assert.NoError(t, service.ComparePassword(hash, "kebadMister"))
	// wrong password
	assert.Error(t, service.ComparePassword(hash, "wrong"))
}
