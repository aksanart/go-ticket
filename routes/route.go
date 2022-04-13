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

	// r.POST("/login", controllers.Login(services.NewUserService()))
	return r
}
