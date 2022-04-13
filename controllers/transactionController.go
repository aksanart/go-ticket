package controllers

import (
	"fmt"

	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/repository"
	"github.com/gin-gonic/gin"
)

func TransactionCheckout(repo repository.TransactionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				var param config.CheckoutParam
				if err := c.ShouldBind(&param); err != nil {
					config.ResponseStruct.Set(4001)
					config.Throw(err.Error())
				}
				err := repo.Save(param)
				if err != nil {
					config.Throw(err.Error())
				}
				config.Response(c, "success checkout", nil)
			},
			Catch: func(e config.Exception) {
				config.ErrorResponse(c, e)
			},
			Finally: func() {
				fmt.Println("selesai")
			},
		}.Do()
	}
}

func TransactionPay(repo repository.TransactionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				var param config.PayParam
				if err := c.ShouldBind(&param); err != nil {
					config.ResponseStruct.Set(4001)
					config.Throw(err.Error())
				}
				err := repo.Pay(param)
				if err != nil {
					config.Throw(err.Error())
				}
				config.Response(c, "success membayar", nil)
			},
			Catch: func(e config.Exception) {
				config.ErrorResponse(c, e)
			},
			Finally: func() {
				fmt.Println("selesai")
			},
		}.Do()
	}
}

func TransactionList(repo repository.TransactionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				username := c.Query("username")
				if username == "" {
					config.ResponseStruct.Set(7001)
					config.Throw("username wajib")
				}
				list, err := repo.FindAllTransactionUser(username)
				if err != nil {
					config.Throw(err.Error())
				}
				config.Response(c, "success list", list)
			},
			Catch: func(e config.Exception) {
				config.ErrorResponse(c, e)
			},
			Finally: func() {
				fmt.Println("selesai CreateEvent")
			},
		}.Do()
	}
}
