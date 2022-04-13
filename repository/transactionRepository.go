package repository

import (
	"github.com/aksanart/go-ticket/config"
)

type TransactionRepository interface {
	FindAll() string
	Save(config.CheckoutParam) error
	Pay(config.PayParam) error
}
