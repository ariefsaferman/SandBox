package handlers

import (
	"log"
	"net/http"

	"gorm-raw/db"
	"gorm-raw/models"

	"github.com/gin-gonic/gin"
)

func FindAllTrainersGorm() (*[]models.Trainer, error) {
	d := db.GetGormDB()
	var trainers []models.Trainer

	err := d.Find(&trainers).Error
	if err != nil {
		return nil, err
	}

	return &trainers, nil
}

func FindAllTrainers() (*[]models.Trainer, error) {
	d := db.GetDB()

	rows, err := d.Query(
		"SELECT * FROM trainers",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var trainers []models.Trainer

	for rows.Next() {
		var trainer models.Trainer
		err := rows.Scan(
			&trainer.ID,
			&trainer.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		trainers = append(trainers, trainer)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}

	return &trainers, err
}

func GetAllTrainers(ctx *gin.Context) {
	var trainers *[]models.Trainer

	trainers, err := FindAllTrainers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    trainers,
		"message": "success",
	})
}

func GetAllTrainersGorm(ctx *gin.Context) {
	var trainers *[]models.Trainer

	trainers, err := FindAllTrainersGorm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    trainers,
		"message": "success",
	})
}

func AddNewTrainer(ctx *gin.Context) {
	d := db.GetDB()
	var trainer models.Trainer

	ctx.BindJSON(&trainer)

	err := d.QueryRow(
		`
			INSERT INTO "trainers" ("id","name") VALUES ($1, $2)
			RETURNING "id", "name";
		`,
		trainer.ID,
		trainer.Name,
	).Scan(
		&trainer.ID,
		&trainer.Name,
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"data":    trainer,
		"message": "Created",
	})
}
