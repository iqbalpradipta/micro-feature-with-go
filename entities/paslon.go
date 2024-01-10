package entities

import "time"

type Paslon struct {
	ID			uint		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	NomorUrut	int			`json:"nomorUrut"`
	VisiMisi	string		`json:"visiMisi"`
	Image 		string		`json:"image"`
	VotePoint 	int			`json:"votePoint"`
	Created_At  time.Time
	Updated_At 	time.Time
}