package main

import (
	"digital-bank/infrastructure/adapter"
	"digital-bank/infrastructure/config"
	"digital-bank/infrastructure/events"
	"digital-bank/infrastructure/http/server"
	"log"
)

// @title Digital Bank API
// @version 1.0
// @description This is a sample server for a digital bank.
// @host localhost:8080
// @BasePath /
func main() {
	log.Println("PDF generado con éxito!")

	config.LoadEnvironmentVariables()

	events.SubscribeToEvents()

	adapter.CreateCustomValidator()

	server.Start()
}
