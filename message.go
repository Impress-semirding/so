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
func listen(channels, reciever string, msgs chan map[string]string) *redis.Message {
	service := map[string]string{
		"reciever": reciever,
		"channels": channels,
	}
	manager := connect()
	pubsub := manager.subscribe(service["channels"])
	defer pubsub.Close()
	for {
		msg, _ := pubsub.ReceiveMessage()
		service["msgBody"] = msg.Payload
		msgs <- service
	}
}
