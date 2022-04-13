package entity

import (
	"context"

	"github.com/aksanart/go-ticket/config"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"email,required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Name     string `json:"name" bson:"name" binding:"required"`
	Gender   string `json:"gender" bson:"gender" binding:"required"`
}

type UserIntf interface {
	FindUserByEmail(email string) (UserEntity, error)
	SaveUser(UserEntity) error
}

func NewUserEntity() *UserEntity {
	return &UserEntity{}
}

func (u *UserEntity) FindUserByEmail(email string) (UserEntity, error) {
	var user UserEntity
	coll := config.Mongo.Conn.Collection("users")
	filter := bson.M{
		"$or": []interface{}{
			bson.M{"email": email},
			bson.M{"username": email},
		},
	}
	err := coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return user, err
		}
	}
	return user, nil
}

func (u *UserEntity) SaveUser(user UserEntity) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	coll := config.Mongo.Conn.Collection("users")
	_, err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}
