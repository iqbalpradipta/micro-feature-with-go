package entities

import (
	"time"

	"gorm.io/gorm"
)



type Articles struct {
	gorm.Model
	ID				uint		`json:"id" gorm:"primaryKey"`
	Title			string		`json:"title"`
	Date			*time.Time	`json:"date"`
	Image			string		`json:"image"`
	Description 	string		`json:"description"`
	Created_At		time.Time
	Updated_At		time.Time
}