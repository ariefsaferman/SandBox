package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=saferman14 dbname=demo_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect DB", err)

	}
	log.Print("db connected")
	db := conn
	return db
}
