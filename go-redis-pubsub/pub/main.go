package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "password",
	DB:       0,
})

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

func main() {
	fmt.Println("tes")
	router := gin.Default()

	router.POST("/user", func(c *gin.Context) {
		body := User{}
		c.ShouldBind(&body)

		payload, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
