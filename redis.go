package main

import (
	"github.com/go-redis/redis"
)

type redisManager struct {
	client *redis.Client
}

func connect() *redisManager {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &redisManager{client}
}

func (server *redisManager) subscribe(channels string) *redis.PubSub {
	return server.client.Subscribe(channels)
}

func (server *redisManager) publish(channel string, message interface{}) {
	server.client.Publish(channel, message)
}

func (server *redisManager) get(key string) string {
	val, err := server.client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

// func ExampleClient() {
// 	err := client.Set("key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := client.Get("key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := client.Get("key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist
// }
