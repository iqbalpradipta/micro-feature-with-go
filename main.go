package main

import (
	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
	"github.com/iqbalpradipta/micro-feature-with-go/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {
	e := echo.New()
	e.Use(middleware.CORS())
	db := config.DBConfig()
	db.AutoMigrate(&entities.Articles{})
	routes.Routes(e)
	
	e.Logger.Fatal(e.Start(":8000"))
}