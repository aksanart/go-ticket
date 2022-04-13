package repository

import "github.com/aksanart/go-ticket/entity"

type EventRepository interface {
	FindAll() string
	Save(entity.EventEntity) error
}
