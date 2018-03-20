package main

import (
	"strconv"

	"github.com/go-redis/redis"
)

type msgService struct {
	senderAddr string
	reciever   string
	channels   string
	msgBody    string
}

type msgMeta struct {
	platform string
	types    string
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
		var s []string
		mm := &msgMeta{}
		mm = msg.Payload
		s = append(s, strconv.Itoa(msg.Payload.platform))
		// s = append(s, strconv.Itoa(msg.Payload.type))
		if service[s] == nil {

		}
		service["msgBody"] = msg.Payload
		msgs <- service
	}
}

func deal(msgs chan map[string]string) {

}
