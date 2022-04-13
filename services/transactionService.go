package services

import (
	"errors"

	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type transactionService struct{}

func NewTransactionService() *transactionService {
	return &transactionService{}
}

func (u *transactionService) FindAll() string {
	return "tes FindAll"
}

func (u *transactionService) Pay(x config.PayParam) error {
	model := entity.NewTransactionEntity()
	filterTrx := bson.M{
		"price_unik": x.PriceUniqueValue,
		"status":     0,
	}
	cek, err := model.FindTransaction(filterTrx)
	if err != nil {
		config.ResponseStruct.Set(5002)
		return err
	}
	if cek.EventName == "" {
		config.ResponseStruct.Set(5003)
		return errors.New("transaction tidak ditemukan")
	}
	if cek.Status > 0 {
		config.ResponseStruct.Set(5004)
		return errors.New("transaction sudah dibayar")
	}
	err = model.UpdateStatus(x.PriceUniqueValue)
	if err != nil {
		config.ResponseStruct.Set(5002)
		return err
	}
	return nil
}

func (u *transactionService) Save(x config.CheckoutParam) error {
	model := entity.NewTransactionEntity()
	filterTrx := bson.M{
		"event_name": x.EventName,
		"username":   x.UserName,
	}
	cek, err := model.FindTransaction(filterTrx)
	if err != nil {
		config.ResponseStruct.Set(4002)
		return err
	}
	if cek.EventName != "" {
		config.ResponseStruct.Set(4003)
		return errors.New("transaction sudah ada")
	}
	eventModel := entity.NewEventEntity()
	cekEvent, err := eventModel.FindEventByName(x.EventName)
	if err != nil {
		config.ResponseStruct.Set(4002)
		return err
	}
	if cekEvent.Ticket < x.TicketCheckout {
		config.ResponseStruct.Set(4003)
		return errors.New("no ticket available")
	}
	userModel := entity.NewUserEntity()
	cekUser, err := userModel.FindUserByEmail(x.UserName)
	if err != nil {
		config.ResponseStruct.Set(4005)
		return err
	}
	if cekUser.Name == "" {
		config.ResponseStruct.Set(4004)
		return errors.New("no user found")
	}
	harga := (cekEvent.Price * float64(x.TicketCheckout))
	unikHarga := harga
	for {
		unikHarga = unikHarga + 1
		filterTrxUnikHarga := bson.M{
			"price_unik": unikHarga,
			"status": bson.M{
				"$ne": 1,
			},
		}
		count, err := model.FindTransactionCount(filterTrxUnikHarga)
		if err != nil {
			config.ResponseStruct.Set(4004)
			return err
		}
		if count == 0 {
			break
		}
	}
	var data entity.TransactionEntity
	data.EventName = x.EventName
	data.UserName = x.UserName
	data.TicketCheckout = x.TicketCheckout
	data.Price = harga
	data.PriceUniqueValue = unikHarga
	data.Status = 0
	err = model.Save(data)
	if err != nil {
		config.ResponseStruct.Set(4004)
		return err
	}
	tiket := cekEvent.Ticket - x.TicketCheckout
	err = eventModel.UpdateTicket(cekEvent, tiket)
	if err != nil {
		config.ResponseStruct.Set(4005)
		return err
	}
	return nil
}
