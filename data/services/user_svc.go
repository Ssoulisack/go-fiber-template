package services

import (
	"go-fiber/data/input"
	"go-fiber/data/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

// GetAllUsers implements UserService.
func (u *UserService) GetAllUsers() {
	panic("unimplemented")
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

var _ input.UserService = (*UserService)(nil)
