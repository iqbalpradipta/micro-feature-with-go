package entities

import "time"

type Users struct {
	ID				uint		`json:"id" gorm:"primaryKey"`
	Fullname		string		`json:"fullName"`
	Alamat			string		`json:"alamat"`
	Jeniskelamin	string		`json:"jenisKelamin"`
	Username 		string		`json:"username"`
	Password 		string		`json:"password"`
	Created_At  	time.Time
	Updated_At 		time.Time	
	Voting			uint
	ArticlesID		[]Articles	`gorm:"foreignKey:Author"`
}