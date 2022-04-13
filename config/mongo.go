package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mongo = &dbMongo{}

type dbMongo struct {
	Conn *mongo.Database
	Err  error
}

func (d *dbMongo) Init() {
	d.Conn, d.Err = d.ConnectDB()
}

func (d *dbMongo) ConnectDB() (*mongo.Database, error) {
	var (
		host       = os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
		user       = os.Getenv("MONGO_USERNAME")
		dbname     = os.Getenv("MONGO_DBNAME")
		password   = os.Getenv("MONGO_PASSWORD")
		authSource = os.Getenv("MONGO_AUTH")
	)
	var connStr string
	if user != "" {
		connStr = fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s&w=majority&readPreference=primary&directConnection=true&ssl=false", user, password, host, authSource)
	} else {
		connStr = fmt.Sprintf("mongodb://%s/?authSource=%s", host, authSource)
	}
	clientOptions := options.Client().ApplyURI(connStr)

	// Connect to MongoDB
	dbTimeout := 5 * time.Second // Default

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection is can access mongo or not
	err = client.Ping(ctx, &readpref.ReadPref{})
	if err != nil {
		log.Fatalf("Failed when ping to: %s, error: %v", clientOptions.GetURI(), err)
	}
	log.Println("Connected to MongoDB:", clientOptions.GetURI())

	return client.Database(dbname), nil
}
