package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dev79844/go-chat/pubsub"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

type Message struct{
    Username    string  `json:"username"`
    Message     string  `json:"message"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func handleConnections(w http.ResponseWriter, r *http.Request){
    conn, err := upgrader.Upgrade(w,r, nil)
    if err!=nil{
        fmt.Println(err)
        return
    }

    defer conn.Close()

    clients[conn] = true

    msgChan := make(chan pubsub.Message)
    go pubsub.Subscribe(msgChan)

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            log.Println(err)
            delete(clients, conn)
            return
        }

        pubsub.Publish(pubsub.Message{
            Username: msg.Username,
            Message:  msg.Message,
        })

        for redisMsg := range msgChan {
            for client := range clients {
                err := client.WriteJSON(Message{
                    Username: redisMsg.Username,
                    Message:  redisMsg.Message,
                })
                if err != nil {
                    log.Println(err)
                    client.Close()
                    delete(clients, client)
                }
            }
        }
    }
}

func handleMessages(){
    for {
     msg := <-broadcast
   
     for client := range clients {
      err := client.WriteJSON(msg)
      if err != nil {
       fmt.Println(err)
       client.Close()
       delete(clients, client)
      }
     }
    }
   }

func StartWebSocketServer() {
    redisClient := pubsub.InitialiseRedis()
    defer redisClient.Close()

    http.HandleFunc("/ws", handleConnections)

    go handleMessages()

    log.Println("Starting WebSocket server on :8081")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}