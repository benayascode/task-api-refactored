package usecases

import (
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"
)

type UserUseCaseImpl struct {
	userRepository  domain.UserRepository
	passwordService infrastructure.PasswordService
	jwtService      infrastructure.JWTService
}

func NewUserUseCase(
	userRepository domain.UserRepository,
	passwordService infrastructure.PasswordService,
	jwtService infrastructure.JWTService,
) domain.UserUseCase {
	return &UserUseCaseImpl{
		userRepository:  userRepository,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (u *UserUseCaseImpl) RegisterUser(username, password string) error {

	hashedPassword, err := u.passwordService.HashPassword(password)
	if err != nil {
		return err
	}

	return u.userRepository.RegisterUser(username, hashedPassword)
}

func (u *UserUseCaseImpl) LoginUser(username, password string) (string, error) {

	user, err := u.userRepository.AuthenticateUser(username, password)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	err = u.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	token, err := u.jwtService.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCaseImpl) PromoteUser(username string) error {
	return u.userRepository.PromoteUser(username)
}
