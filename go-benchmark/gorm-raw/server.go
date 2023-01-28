package main

import (
	"gorm-raw/db"
	"gorm-raw/handlers"
	"log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	db.ConnectDB()
	db.ConnectGormDB()

	r := gin.Default()
	pprof.Register(r)

	v1 := r.Group("/v1")
	{
		v1.GET("/trainers", handlers.GetAllTrainers)
		v1.POST("/trainers", handlers.AddNewTrainer)
		v1.POST("/register", handlers.AddNewUser)
		v1.POST("/login", handlers.Login)

		gorm := v1.Group("/gorm")
		{
			gorm.GET("/trainers", handlers.GetAllTrainersGorm)
			gorm.POST("/login", handlers.LoginGorm)
		}
	}

	r.Run(":3000")
}
