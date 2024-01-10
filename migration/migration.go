package migration

import (
	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
)

func RunMigration()  {
	err := config.DB.AutoMigrate(
		&entities.Articles{},
	)

	if err != nil {
		panic(err)
	}
}