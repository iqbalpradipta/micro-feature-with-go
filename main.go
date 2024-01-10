package main

import (
	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/iqbalpradipta/micro-feature-with-go/migration"
	"github.com/iqbalpradipta/micro-feature-with-go/routes"
	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()

	config.DBConfig()
	migration.RunMigration()
	routes.Routes(e.Group("/api/v1"))
	
	e.Logger.Fatal(e.Start(":8000"))
}