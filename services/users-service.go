package service

import (
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
	"gorm.io/gorm"
)

type UsersService interface {
	GetUsers() ([]entities.Users, error)
	GetUsersById(id int) (entities.Users, error)
}

func RepositoryUsers(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetUsers() ([]entities.Users, error) {
	var users []entities.Users
	err := r.db.Preload("ArticlesID").Find(&users).Error

	return users, err
}

func (r *repository) GetUsersById(id int) (entities.Users, error) {
	var users entities.Users
	err := r.db.Preload("ArticlesID").First(&users, id).Error

	return users, err
}