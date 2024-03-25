package service

import (
	"github.com/cho8833/CC-Calendar/internal/user/model"
	"github.com/cho8833/CC-Calendar/internal/user/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (svc UserService) GetUser(userId int64) (*model.User, error) {
	user, err := svc.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}