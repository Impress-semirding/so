package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

// need struct msg.
func messageManager() *redis.Message {
	channels := "btc-msg"
	manager := connect()
	pubsub := manager.subscribe(channels)
	defer pubsub.Close()
	for {
		msg, _ := pubsub.ReceiveMessage()
		fmt.Println(msg.Channel, msg.Payload)
		// return msg
	}
}

func main() {
	go messageManager()

	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
