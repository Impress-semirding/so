package main

import (
	"github.com/go-redis/redis"
)

type msgService struct {
	senderAddr string
	reciever   string
	channels   string
	msgBody    string
}

// need struct msg.
func listen(channels, reciever string, result chan map[string]string) *redis.Message {
	// service := &msgService{
	// 	reciever: reciever,
	// 	channels: channels,
	// }
	service := map[string]string{
		reciever: "localhosyt",
		channels: channels,
	}
	manager := connect()
	pubsub := manager.subscribe(service["channels"])
	defer pubsub.Close()
	for {
		msg, _ := pubsub.ReceiveMessage()
		service["msgBody"] = msg.Payload
		result <- service
	}
}
