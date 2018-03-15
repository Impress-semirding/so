package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
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

func main() {
	channels := "dingxue"
	manager := connect()
	pubsub := manager.subscribe(channels)
	defer pubsub.Close()
	subscr, err1 := pubsub.ReceiveTimeout(time.Second)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(subscr)
	me := &message{id: 12121}
	manager.publish(channels, me.id)
	msg, err2 := pubsub.ReceiveMessage()
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(msg.Channel, msg.Payload)

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
