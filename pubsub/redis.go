package pubsub

import (
	"context"
	"log"
	"os"
	"strings"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
    redisClient *redis.Client
    redisChannel = "chat"
)

type Message struct {
    Username string
    Message  string
}

func InitialiseRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr: 		os.Getenv("REDIS_CONNECTION_STRING"),
		Password: 	"",
		DB:	  		0,
	})

	ping, err := conn.Ping(context.Background()).Result()
	if err != nil{
		log.Fatal("Error connecting to redis", err)
	}

	log.Println("Connected to redis", ping)

	redisClient = conn

	return redisClient
}

func Subscribe(ch chan <- Message){
	pubsub := redisClient.Subscribe(context.Background(), redisChannel)

	defer pubsub.Close()

	for{
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
            log.Println(err)
            continue
        }

		if msg.Channel == redisChannel{
			ch <- Message{
                Username: string(strings.Split(msg.Payload, ":")[0]),
                Message:  string(strings.Split(msg.Payload, ":")[1]),
            }
		}
	}
}

func Publish(msg Message) {
    payload := fmt.Sprintf("%s:%s", msg.Username, msg.Message)
    redisClient.Publish(context.Background(), redisChannel, payload)
}