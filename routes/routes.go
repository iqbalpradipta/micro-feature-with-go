package routes

import (
	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/iqbalpradipta/micro-feature-with-go/controller"
	service "github.com/iqbalpradipta/micro-feature-with-go/services"
	"github.com/labstack/echo/v4"
)


func Routes(e *echo.Group) {

	// Article
	ra := service.RepositoryArticle(config.DB)
	cra := controller.ArticlesController(ra)

	e.GET("/articles", cra.GetArticle)
	e.GET("/articles/:id", cra.GetArticleById)

	// Paslon
	rp := service.RepositoryPaslon(config.DB)
	crp := controller.PaslonController(rp)

	e.GET("/paslon", crp.GetPaslon)
	e.GET("/paslon/:id", crp.GetPaslonById)


}