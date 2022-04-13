package controllers

import (
	"fmt"

	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/entity"
	"github.com/aksanart/go-ticket/repository"
	"github.com/gin-gonic/gin"
)

func CreateUser(repo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				var user entity.UserEntity
				if err := c.ShouldBind(&user); err != nil {
					config.ResponseStruct.Set(2001)
					config.Throw(err.Error())
				}
				err := repo.Save(user)
				if err != nil {
					config.Throw(err.Error())
				}
				config.Response(c, "success register user", nil)
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

func UserList(repo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				username := c.Query("username")
				if username == "" {
					config.ResponseStruct.Set(5001)
					config.Throw("username wajib diisi")
				}
				user, err := repo.FindUserByEmail(username)
				if err != nil {
					config.Throw(err.Error())
				}
				config.Response(c, "success register user", user)
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
