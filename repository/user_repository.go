package repository

import (
	"gorestapi/model"

	"gorm.io/gorm"
)

type IUserReposiroty interface {
	FindByEmail(user *model.User, email string) error
	Create(user *model.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserReposiroty {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) FindByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Create(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
