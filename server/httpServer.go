package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	handler "github.com/Dev79844/go-chat/handlers"
)

func StartHTTPServer(){
	app := gin.Default()

	app.GET("/health",handler.Health)

	err := app.Run(os.Getenv("HTTP_PORT"))
	if err!=nil{
		log.Fatal("Error starting the http server", err)
	}

	log.Println("HTTP server started on port ",os.Getenv("HTTP_PORT"))
}