package usecase

import (
	"errors"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/repository"
	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
)

type UserUsecase interface {
	Login(user *entity.UserLoginRequest) (string, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

type UserConfig struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(cfg *UserConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository: cfg.UserRepository,
	}
}

func (u *userUsecaseImpl) Login(user *entity.UserLoginRequest) (string, error) {
	searchedUser, err := u.userRepository.GetOneUserByEmail(user.Email)
	if err != nil {
		return "", errors.New("user is not found")
	}

	isMatch := util.ComparePassword(searchedUser.Password, user.Password)
	if !isMatch {
		return "", errors.New("invalid password")
	}

	token, err := util.GenerateAccessToken(searchedUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
