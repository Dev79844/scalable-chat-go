package main

import (
	"log"

	"github.com/Dev79844/go-chat/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("unable to import .env file")
	}
}

func main(){
	// server.StartHTTPServer()
	server.StartWebSocketServer()
}