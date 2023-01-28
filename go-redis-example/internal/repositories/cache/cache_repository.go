package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"git.garena.com/aldino.rahman/go-redis-example/internal/dto/request"
	"git.garena.com/aldino.rahman/go-redis-example/internal/models"
	"git.garena.com/aldino.rahman/go-redis-example/internal/repositories"
	"github.com/go-redis/redis/v9"
)



type studentRepository struct {
	rdc *redis.Client
}

func NewStudentCacheRepository(rdc *redis.Client) repositories.ICacheRepository {
	return &studentRepository{
		rdc:rdc,
	}
}

var ctx = context.Background()

func (sr *studentRepository) GetStudentCache(key string) (*models.Student,error)  {
	// using list
	val,err := sr.rdc.Get(ctx,key).Result()
	if err == redis.Nil {
		// key does not exist
        return nil, fmt.Errorf("GetData:%s does not exist",key) 
    } else if err != nil {
       return nil,fmt.Errorf("Getdata : %w",err)
    }
	var data *models.Student
	err = json.Unmarshal([]byte(val), &data)

	return data,nil
}

func (sr *studentRepository) SetStudentCache(key string,studentData *request.StudentDTO,expiry time.Duration) error {
	// set list data type
	jsonData,_:= json.Marshal(studentData)
	err:= sr.rdc.Set(ctx,key,string(jsonData),expiry).Err()
	if err != nil{
		return err
	}
	return nil
}