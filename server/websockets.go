package server

import(
	"net/http"

	redis "github.com/Dev79844/go-chat/pubsub"
)

func StartWebSocketServer() {
	redisClient := redis.InitialiseRedis()
	defer redisClient.Close()

	http.ListenAndServe(":8081", nil)
}