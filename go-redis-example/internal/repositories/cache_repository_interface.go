package repositories

import (
	"time"

	"git.garena.com/aldino.rahman/go-redis-example/internal/dto/request"
	"git.garena.com/aldino.rahman/go-redis-example/internal/models"
)

type ICacheRepository interface {
	GetStudentCache(key string) (*models.Student,error)
	SetStudentCache(key string,studentData *request.StudentDTO,expiry time.Duration) (error)
}