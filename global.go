package main

type order struct {
	a int
	b int
}

var coincola = make(map[string][]map[string]interface{})
var coincola_btch interface{}

var redisClient *redisManager
