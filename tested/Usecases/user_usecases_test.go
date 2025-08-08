package usecases

import (
	"errors"
	"testing"

	domain "task-manager/Domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUseCase_RegisterUser_PersistsHashed(t *testing.T) {
	repo := new(MockUserRepository)
	pass := new(MockPasswordService)
	jwt := new(MockJWTService)
	uc := NewUserUseCase(repo, pass, jwt)

	pass.On("HashPassword", "plain").Return("hashed", nil).Once()
	repo.On("RegisterUser", "bob", "hashed").Return(nil).Once()

	err := uc.RegisterUser("bob", "plain")
	assert.NoError(t, err)
	pass.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestUserUseCase_LoginUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	pass := new(MockPasswordService)
	jwt := new(MockJWTService)
	uc := NewUserUseCase(repo, pass, jwt)

	storedUser := domain.User{UserName: "bob", Password: "hashed", Role: "user"}
	repo.On("AuthenticateUser", "bob", mock.Anything).Return(storedUser, nil).Once()
	pass.On("ComparePassword", "hashed", "plain").Return(nil).Once()
	jwt.On("GenerateToken", storedUser).Return("token123", nil).Once()

	token, err := uc.LoginUser("bob", "plain")
	assert.NoError(t, err)
	assert.Equal(t, "token123", token)
	repo.AssertExpectations(t)
	pass.AssertExpectations(t)
	jwt.AssertExpectations(t)
}

func TestUserUseCase_LoginUser_InvalidPassword(t *testing.T) {
	repo := new(MockUserRepository)
	pass := new(MockPasswordService)
	jwt := new(MockJWTService)
	uc := NewUserUseCase(repo, pass, jwt)

	storedUser := domain.User{UserName: "bob", Password: "hashed", Role: "user"}
	repo.On("AuthenticateUser", "bob", mock.Anything).Return(storedUser, nil).Once()
	pass.On("ComparePassword", "hashed", "wrong").Return(errors.New("mismatch")).Once()

	_, err := uc.LoginUser("bob", "wrong")
	assert.ErrorIs(t, err, domain.ErrInvalidCredentials)
	jwt.AssertNotCalled(t, "GenerateToken", mock.Anything)
}

func TestUserUseCase_LoginUser_UserNotFound(t *testing.T) {
	repo := new(MockUserRepository)
	pass := new(MockPasswordService)
	jwt := new(MockJWTService)
	uc := NewUserUseCase(repo, pass, jwt)

	repo.On("AuthenticateUser", "alice", mock.Anything).Return(domain.User{}, errors.New("not found")).Once()

	_, err := uc.LoginUser("alice", "whatever")
	assert.ErrorIs(t, err, domain.ErrInvalidCredentials)
	pass.AssertNotCalled(t, "ComparePassword", mock.Anything, mock.Anything)
	jwt.AssertNotCalled(t, "GenerateToken", mock.Anything)
}

func TestUserUseCase_PromoteUser(t *testing.T) {
	repo := new(MockUserRepository)
	pass := new(MockPasswordService)
	jwt := new(MockJWTService)
	uc := NewUserUseCase(repo, pass, jwt)

	repo.On("PromoteUser", "bob").Return(nil).Once()
	assert.NoError(t, uc.PromoteUser("bob"))
	repo.AssertExpectations(t)
}
