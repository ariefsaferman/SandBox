package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "password",
	DB:       0,
})

func main() {

	subscriber := redisClient.Subscribe(ctx, "send-user-data")

	user := User{}
	numb := 0
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		numb++
		fmt.Println(numb)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
	}

}
