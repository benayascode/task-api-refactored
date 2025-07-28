package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
	"task_manager/Domain"
)

type BcryptPasswordService struct{}

func NewBcryptPasswordService() Domain.PasswordService {
	return &BcryptPasswordService{}
}

func (s *BcryptPasswordService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (s *BcryptPasswordService) CheckPasswordHash(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
