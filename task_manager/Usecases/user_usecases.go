package Usecases

import (
	"context"
	"task_manager/Domain"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
)

type UserUseCase struct {
	UserRepo *Repositories.UserRepository
}

func (uu *UserUseCase) Register(ctx context.Context, username, password string) error {
	hashed, err := Infrastructure.HashPassword(password)
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
	if !Infrastructure.CheckPasswordHash(user.Password, password) {
		return "", err
	}
	return Infrastructure.GenerateToken(user)
}

func (uu *UserUseCase) Promote(ctx context.Context, username string) error {
	return uu.UserRepo.Promote(ctx, username)
}
