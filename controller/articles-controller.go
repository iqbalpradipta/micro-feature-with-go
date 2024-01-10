package controller

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/micro-feature-with-go/helpers"
	service "github.com/iqbalpradipta/micro-feature-with-go/services"
	"github.com/labstack/echo/v4"
)

type controllerArticle struct{
	ArticleRepository service.ArticleService
}

func ArticlesController(articleServices service.ArticleService) *controllerArticle {
	return &controllerArticle{articleServices}
}

func (cr *controllerArticle) GetArticle(c echo.Context) error {
	articles, err := cr.ArticleRepository.GetArticle()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get article Success", articles))
}

func (cr *controllerArticle) GetArticleById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	articles, err := cr.ArticleRepository.GetArticleById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get article Success", articles))
}