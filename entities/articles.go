package entities

import "time"

type Articles struct {
	ID				uint		`gorm:"primaryKey"`
	title			string
	date			*time.Time
	image			string
	description 	string
}