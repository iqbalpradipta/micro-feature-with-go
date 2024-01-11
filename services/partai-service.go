package service

import (
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
	"gorm.io/gorm"
)

type PartaiService interface {
	GetPartai() ([]entities.Partai, error)
	GetPartaiById(id int) (entities.Partai, error)
}

func RepositoryPartai(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetPartai() ([]entities.Partai, error) {
	var partai []entities.Partai
	err := r.db.Find(&partai).Error

	return partai, err
}

func (r *repository) GetPartaiById(id int) (entities.Partai, error) {
	var partai entities.Partai
	err := r.db.First(&partai, id).Error

	return partai, err
}