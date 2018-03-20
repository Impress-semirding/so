package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type msgService struct {
	senderAddr string
	reciever   string
	channels   string
	msgBody    string
}

// type Message struct {
// 	Asks []Order `json:"Bids"`
// 	Bids []Order `json:"Asks"`
// }

// type Order struct {
// 	Price  float64
// 	Volume float64
// }
type msgMeta struct {
	platform string
	types    string
}

// need struct msg.
func listen(channels, reciever string, msgs chan string) *redis.Message {
	service := map[string]string{
		"reciever": reciever,
		"channels": channels,
	}
	manager := connect()
	pubsub := manager.subscribe(service["channels"])
	defer pubsub.Close()
	for {
		msg, _ := pubsub.ReceiveMessage()
		// var data []byte = []byte(msg.Payload)
		var data map[string]string
		if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
			fmt.Println(err)
		}
		sinal := data["platform"] + "-" + data["types"]
		fmt.Print(sinal)
		// service["msgBody"] = msg.Payload
		msgs <- sinal
	}
}

func recieveSinal(msgs chan string) {

}
