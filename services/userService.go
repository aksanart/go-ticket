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

func (u *userService) FindUserByEmail(x string) (user1 entity.UserEntity, err error) {
	model := entity.NewUserEntity()
	cekUser, err := model.FindUserByEmail(x)
	if err != nil {
		config.ResponseStruct.Set(5002)
		return user1, err
	}
	if cekUser.Email == "" {
		config.ResponseStruct.Set(5003)
		return user1, errors.New("user not found")
	}
	return cekUser, nil
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
