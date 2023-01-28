package config

import "github.com/go-redis/redis/v9"



var rdc *redis.Client

func Get_Redis() *redis.Client {
	rdc = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "password",
		DB: 0,
	})
	return rdc
}