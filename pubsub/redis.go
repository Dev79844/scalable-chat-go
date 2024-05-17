package pubsub

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

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