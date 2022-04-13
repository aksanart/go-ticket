package controllers

import (
	"fmt"

	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/entity"
	"github.com/aksanart/go-ticket/repository"
	"github.com/gin-gonic/gin"
)

func CreateEvent(repo repository.EventRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				var data entity.EventEntity
				if err := c.ShouldBind(&data); err != nil {
					config.ResponseStruct.Set(3001)
					config.Throw(err.Error())
				}
				err := repo.Save(data)
				if err != nil {
					config.Throw(err.Error())
				}
				config.Response(c, "success save event", nil)
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
