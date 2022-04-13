package entity

import (
	"context"

	"github.com/aksanart/go-ticket/config"
	"go.mongodb.org/mongo-driver/bson"
)

type EventEntity struct {
	Name   string  `json:"name" bson:"name" binding:"required"`
	Ticket int32   `json:"ticket_available" bson:"ticket_available" binding:"required,numeric"`
	Price  float64 `json:"Price" bson:"Price" binding:"required"`
}

type EventIntf interface {
	FindEventByName(x string) (EventEntity, error)
	Save(EventEntity) error
	Update(EventEntity) error
	FindEvent(e string) ([]EventEntity, error)
}

func NewEventEntity() *EventEntity {
	return &EventEntity{}
}

func (u *EventEntity) FindEventByName(e string) (EventEntity, error) {
	var data EventEntity
	coll := config.Mongo.Conn.Collection("event")
	filter := bson.M{"name": e}
	err := coll.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return data, err
		}
	}
	return data, nil
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
	filter := bson.M{"name": e.Name}
	update := bson.M{"$set": bson.M{"ticket_available": tiket}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (u *EventEntity) FindEvent() ([]EventEntity, error) {
	var data []EventEntity
	coll := config.Mongo.Conn.Collection("event")
	filter := bson.M{}
	curs, err := coll.Find(context.Background(), filter)
	if err != nil {
		return data, err
	}
	defer curs.Close(context.Background())
	for curs.Next(context.Background()) {
		var singleRow EventEntity
		err := curs.Decode(&singleRow)
		if err != nil {
			return data, err
		}
		data = append(data, singleRow)
	}
	return data, nil
}
