package main

import (
	"log"
	"os"

	"github.com/aksanart/go-ticket/config"
	"github.com/aksanart/go-ticket/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	config.Mongo.Init()
	r := routes.SetupRoutes()
	r.Run(os.Getenv("SERVICE_PORT"))
}
