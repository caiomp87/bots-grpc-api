package main

import (
	"log"

	"github.com/caiomp87/bots-grpc-api/database"
	"github.com/caiomp87/bots-grpc-api/repository"
	"github.com/caiomp87/bots-grpc-api/services"
)

func main() {
	client, err := database.Connect()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	collection := client.Database("poc").Collection("bot")
	repository.BotCollection = repository.InitBot(collection)

	services.StartServer()
}
