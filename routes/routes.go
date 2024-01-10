package routes

import (
	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/iqbalpradipta/micro-feature-with-go/controller"
	service "github.com/iqbalpradipta/micro-feature-with-go/services"
	"github.com/labstack/echo/v4"
)


func Routes(e *echo.Group) {

	r := service.RepositoryArticle(config.DB)
	cr := controller.ArticlesController(r)

	e.GET("/articles", cr.GetArticle)
}