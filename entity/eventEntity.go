package entity

import (
	"context"

	"github.com/aksanart/go-ticket/config"
	"go.mongodb.org/mongo-driver/bson"
)

type EventEntity struct {
	Name   string `json:"name" bson:"name" binding:"required"`
	Ticket int32  `json:"ticket_available" bson:"ticket_available" binding:"required,numeric"`
}

type EventIntf interface {
	FindEventByName(x string) (EventEntity, error)
	Save(EventEntity) error
	Update(EventEntity) error
}

func NewEventEntity() *EventEntity {
	return &EventEntity{}
}

func (u *EventEntity) FindEventByName(e string) (EventEntity, error) {
	var user EventEntity
	coll := config.Mongo.Conn.Collection("event")
	filter := bson.M{"name": e}
	err := coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return user, err
		}
	}
	return user, nil
}

func (u *EventEntity) Save(e EventEntity) error {
	coll := config.Mongo.Conn.Collection("event")
	_, err := coll.InsertOne(context.TODO(), e)
	if err != nil {
		return err
	}
	return nil
}

func (u *EventEntity) UpdateTicket(e EventEntity, tiket int32) error {
	coll := config.Mongo.Conn.Collection("event")
	filter := bson.M{"name": e}
	update := bson.M{"$set": bson.M{"ticket_available": tiket}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
