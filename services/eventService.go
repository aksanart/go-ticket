package services

import (
	"errors"

	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/entity"
)

type eventService struct{}

func NewEventService() *eventService {
	return &eventService{}
}

func (u *eventService) FindAll() string {
	return "tes FindAll"
}

func (u *eventService) Save(x entity.EventEntity) error {
	model := entity.NewEventEntity()
	cek, err := model.FindEventByName(x.Name)
	if err != nil {
		config.ResponseStruct.Set(3002)
		return err
	}
	if cek.Name != "" {
		config.ResponseStruct.Set(3003)
		return errors.New("event already registered")
	}
	err = model.Save(x)
	if err != nil {
		return nil
	}
	return nil
}
