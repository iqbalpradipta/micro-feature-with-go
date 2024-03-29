package entities

import (
	"time"
)

type Articles struct {
	ID				uint		`json:"id" gorm:"primaryKey"`
	Title			string		`json:"title"`
	Image			string		`json:"image"`
	Description 	string		`json:"description"`
	Created_At		time.Time
	Updated_At		time.Time
	Author			uint
}