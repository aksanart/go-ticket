package entity

import (
	"context"

	"github.com/aksanart/go-ticket/config"
	"go.mongodb.org/mongo-driver/bson"
)

type TransactionEntity struct {
	EventName        string  `json:"event_name" bson:"event_name" binding:"required"`
	UserName         string  `json:"username" bson:"username" binding:"required"`
	TicketCheckout   int32   `json:"ticket_checkout" bson:"ticket_checkout" binding:"required,numeric,min=1"`
	Price            float64 `json:"price" bson:"price" binding:"required"`
	PriceUniqueValue float64 `json:"price_unik" bson:"price_unik" binding:"required"`
	Status           int32   `json:"status" bson:"status" binding:"required"` // 0 checkout, 1 paid
}

type TransactionIntf interface {
	Save(TransactionEntity) error
	FindTransaction(filter interface{}) (TransactionEntity, error)
	FindTransactionCount(filter interface{}) (int64, error)
	UpdateStatus(e float64) error
}

func NewTransactionEntity() *TransactionEntity {
	return &TransactionEntity{}
}

func (u *TransactionEntity) Save(data TransactionEntity) error {
	coll := config.Mongo.Conn.Collection("transaction")
	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (u *TransactionEntity) FindTransaction(filter interface{}) (TransactionEntity, error) {
	var data TransactionEntity
	coll := config.Mongo.Conn.Collection("transaction")
	err := coll.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return data, err
		}
	}
	return data, nil
}

func (u *TransactionEntity) FindTransactionCount(filter interface{}) (int64, error) {
	var data int64
	coll := config.Mongo.Conn.Collection("transaction")
	data, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (u *TransactionEntity) UpdateStatus(e float64) error {
	coll := config.Mongo.Conn.Collection("event")
	filter := bson.M{"price_unik": e}
	update := bson.M{"$set": bson.M{"status": 1}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
