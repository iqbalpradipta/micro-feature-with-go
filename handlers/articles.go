package handlers

import (
	"net/http"

	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/iqbalpradipta/micro-feature-with-go/entities"
	"github.com/labstack/echo/v4"
)

func GetArticle(c echo.Context) error {
	config := config.DBConfig()
	data := config.First(&entities.Articles{})
	return c.JSON(http.StatusOK, data)
}