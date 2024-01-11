package service

import (
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
	"gorm.io/gorm"
)

type PaslonService interface {
	GetPaslon() ([]entities.Paslon, error)
	GetPaslonById(id int) (entities.Paslon, error)
}

func RepositoryPaslon(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetPaslon() ([]entities.Paslon, error) {
	var paslon []entities.Paslon
	err := r.db.Preload("Koalisi").Find(&paslon).Error
	
	return paslon, err
}

func (r *repository) GetPaslonById(id int) (entities.Paslon, error) {
	var paslon entities.Paslon
	err := r.db.Preload("Koalisi").First(&paslon, id).Error

	return paslon, err
}