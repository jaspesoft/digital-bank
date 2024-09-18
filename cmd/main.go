package main

import (
	"digital-bank/infrastructure/api/server"
)

// @title Digital Bank API
// @version 1.0
// @description This is a sample server for a digital bank.
// @host localhost:8080
// @BasePath /
func main() {
	server.Start()
}
