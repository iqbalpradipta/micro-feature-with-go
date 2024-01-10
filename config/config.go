package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConfig() *gorm.DB  {
	connection := "host=localhost user=postgres password=mbangg12 dbname=belajar-go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	
	return db
}
