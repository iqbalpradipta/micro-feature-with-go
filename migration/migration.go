package migration

import (
	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
)

func RunMigration()  {
	err := config.DB.AutoMigrate(
		&entities.Articles{},
		&entities.Paslon{},
		&entities.Partai{},
		&entities.Users{},
		&entities.Votters{},
	)

	if err != nil {
		panic(err)
	}
}