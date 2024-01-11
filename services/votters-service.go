package service

import (
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
	"gorm.io/gorm"
)

type VottersService interface {
	GetVotters() ([]entities.Votters, error)
	GetVottersById(id int) (entities.Votters, error)
}

func RepositoryVotters(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetVotters() ([]entities.Votters, error) {
	var votters []entities.Votters
	err := r.db.Preload("Paslon").Preload("Users").Find(&votters).Error
	
	return votters, err
}

func (r *repository) GetVottersById(id int) (entities.Votters, error) {
	var votters entities.Votters
	err := r.db.Preload("Paslon").Preload("Users").First(&votters, id).Error

	return votters, err
}