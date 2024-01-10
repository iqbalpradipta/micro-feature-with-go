package routes

import (
	"github.com/iqbalpradipta/micro-feature-with-go/handlers"

	"github.com/labstack/echo/v4"
)


func Routes(e *echo.Echo) {
	e.GET("/articles", handlers.GetArticle)
}