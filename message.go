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
func listen(channels, reciever string, msgs chan map[string]string) *redis.Message {
	service := map[string]string{
		"reciever": reciever,
		"channels": channels,
	}
	manager := connect()
	redisClient = manager
	pubsub := manager.subscribe(service["channels"])
	defer pubsub.Close()
	for {
		msg, _ := pubsub.ReceiveMessage()
		// var data []byte = []byte(msg.Payload)
		var data map[string]string
		if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
			fmt.Println(err)
		}
		// sinal := data["platform"] + "-" + data["types"]
		// service["msgBody"] = msg.Payload
		msgs <- data
	}
}

func recieveSinal(sinals chan map[string]string) {
	for range sinals {
		s := <-sinals
		switch {
		case s["status"] == "init":
			key := s["platform"] + "_" + s["types"]

			list := redisClient.get(key)
			var test []map[string]interface{}
			if err := json.Unmarshal([]byte(list), &test); err == nil {

			} else {
				fmt.Println(err)
			}
			coincola[key] = test
		case s["status"] == "update":
			fmt.Print(111)
		default:
			fmt.Print(222)
		}
	}
}
