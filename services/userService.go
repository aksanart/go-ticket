package services

import (
	"errors"

	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/entity"
)

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) FindAll() string {
	return "tes FindAll"
}

func (u *userService) Save(user entity.UserEntity) error {
	model := entity.NewUserEntity()
	cekUser, err := model.FindUserByEmail(user.Email)
	if err != nil {
		config.ResponseStruct.Set(2002)
		return err
	}
	if cekUser.Email != "" {
		config.ResponseStruct.Set(2003)
		return errors.New("user already registered")
	}
	err = model.SaveUser(user)
	if err != nil {
		return nil
	}
	return nil
}
