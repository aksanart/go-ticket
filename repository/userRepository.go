package repository

import "github.com/aksanart/go-ticket/entity"

type UserRepository interface {
	FindAll() string
	Save(entity.UserEntity) error
}
