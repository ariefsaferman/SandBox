package usecases

import (
	"time"

	"git.garena.com/aldino.rahman/go-redis-example/internal/dto/request"
	"git.garena.com/aldino.rahman/go-redis-example/internal/models"
	"git.garena.com/aldino.rahman/go-redis-example/internal/repositories"
)

type IstudentUseCase interface {
	GetStudentbyId(id string)(*models.Student,error) // should be model 
	AddStudent(*request.StudentDTO) error
}


type OptsStudentUseCase struct {
	CacheRepository repositories.ICacheRepository
}

type studentUseCase struct {
	cacheRepository repositories.ICacheRepository
}

func NewStudentUseCase(osu *OptsStudentUseCase) IstudentUseCase{
	return &studentUseCase{
		cacheRepository: osu.CacheRepository,
		// you can add another repo to another database
	}
}

func (s *studentUseCase) GetStudentbyId(id string) (*models.Student,error)  {
	studentData,_:= s.cacheRepository.GetStudentCache(id)
	
	return studentData,nil
}

func (s *studentUseCase) AddStudent(studentData *request.StudentDTO) error {
	// pretend id from other database
	id:="102b"
	// set data to another database
	//  set json data for 1 day
	err:=s.cacheRepository.SetStudentCache(id,studentData,60*time.Second)
	return err
}


