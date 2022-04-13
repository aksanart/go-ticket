package repository

import (
	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/entity"
)

type TransactionRepository interface {
	FindAll() string
	Save(config.CheckoutParam) error
	Pay(config.PayParam) error
	FindAllTransactionUser(username string) ([]entity.TransactionEntity, error)
}
