package entities

type Paslon struct {
	ID			uint		`gorm:"primaryKey"`
	name		string
	nomorUrut	int
	visiMisi	string
	image 		string
	votePoint 	int
}