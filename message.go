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

func dealMsg(s map[string]string) *[]map[string]interface{} {
	key := s["platform"] + "_" + s["types"]

	list := redisClient.get(key)
	var values *[]map[string]interface{}
	// var values = make([]map[string]interface{}, 10, 50)
	if err := json.Unmarshal([]byte(list), &values); err == nil {

	} else {
		fmt.Println(err)
	}
	return values
}

func recieveSinal(sinals chan map[string]string) {
	for {
		s := <-sinals
		switch {
		case s["status"] == "init":
			msg := dealMsg(s)
			//	init, put redis data on local storage.
			for _, num := range *msg {
				fmt.Print(num["a"])
			}

		case s["status"] == "update":
			msg := dealMsg(s)
			// update, if some data change,compare it with local storage other data.
			for _, num := range *msg {
				fmt.Print(num["a"])
			}

		default:
			fmt.Print(222)
		}
	}
}
