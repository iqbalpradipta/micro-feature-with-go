package entities

import "time"

type Votters struct {
	ID				uint		`json:"id" gorm:"primaryKey"`
	Created_At  	time.Time
	Updated_At 		time.Time	
	Paslon			[]Paslon	`gorm:"foreignKey:Voting"`
	Users			[]Users		`gorm:"foreignKey:Voting"`
}