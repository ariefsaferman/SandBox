package repository

import (
	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetOneUserByEmail(e string) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserConfig) UserRepository {
	return &userRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *userRepositoryImpl) GetOneUserByEmail(e string) (*entity.User, error) {
	var res *entity.User
	err := r.db.Where("email = ?", e).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
