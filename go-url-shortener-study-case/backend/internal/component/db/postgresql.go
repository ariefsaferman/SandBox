package db

import (
	"fmt"
	"log"
	"url-shortener/internal/component/config"
	"url-shortener/internal/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresSQLDB(cfg config.Config) *gorm.DB {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Port,
	)

	var loggerOption logger.Interface

	if cfg.Application.Environment == constants.EnvironmentProduction {
		loggerOption = logger.Default.LogMode(logger.Error)
	} else {
		loggerOption = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerOption,
	})
	if err != nil {
		log.Fatal("failed to open connection: ", err)
	}

	return db
}
