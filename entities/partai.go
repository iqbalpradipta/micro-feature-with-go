package entities

import "time"

type Partai struct {
	ID			uint		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	Ketum		string		`json:"ketum"`
	VisiMisi	string		`json:"visiMisi"`
	Alamat 		string		`json:"alamat"`
	Image 		string		`json:"image"`
	Created_At  time.Time
	Updated_At 	time.Time
	Paslon		uint
}