package Usecases

import (
	"context"
	"task_manager/Domain"
)

type UserUseCase struct {
	UserRepo        Domain.UserRepository
	PasswordService Domain.PasswordService
	JWTService      Domain.JWTService
}

func NewUserUseCase(userRepo Domain.UserRepository, passwordService Domain.PasswordService, jwtService Domain.JWTService) *UserUseCase {
	return &UserUseCase{
		UserRepo:        userRepo,
		PasswordService: passwordService,
		JWTService:      jwtService,
	}
}

func (uu *UserUseCase) Register(ctx context.Context, username, password string) error {
	hashed, err := uu.PasswordService.HashPassword(password)
	if err != nil {
		return err
	}
	user := Domain.User{
		UserName: username,
		Password: hashed,
		Role:     "user",
	}
	return uu.UserRepo.Register(ctx, user)
}

func (uu *UserUseCase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := uu.UserRepo.Authenticate(ctx, username)
	if err != nil {
		return "", err
	}
	if !uu.PasswordService.CheckPasswordHash(user.Password, password) {
		return "", err
	}
	return uu.JWTService.GenerateToken(user)
}

func (uu *UserUseCase) Promote(ctx context.Context, username string) error {
	return uu.UserRepo.Promote(ctx, username)
}
