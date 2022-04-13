package routes

import (
	"github.com/aksanart/go-ticket/controllers"
	"github.com/aksanart/go-ticket/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.POST("users/register", controllers.CreateUser(services.NewUserService()))
	r.POST("event/create", controllers.CreateEvent(services.NewEventService()))
	r.POST("transaction/checkout", controllers.TransactionCheckout(services.NewTransactionService()))
	r.POST("transaction/bayar", controllers.TransactionPay(services.NewTransactionService()))
	r.GET("users/list", controllers.UserList(services.NewUserService()))
	r.GET("event/list", controllers.EventList(services.NewEventService()))
	r.GET("transaction/list", controllers.TransactionList(services.NewTransactionService()))

	return r
}
