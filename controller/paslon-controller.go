package controller

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/micro-feature-with-go/helpers"
	service "github.com/iqbalpradipta/micro-feature-with-go/services"
	"github.com/labstack/echo/v4"
)

type controllerPaslon struct {
	PaslonRepository service.PaslonService
}

func PaslonController(paslonService service.PaslonService) *controllerPaslon {
	return &controllerPaslon{paslonService}
}

func (cr *controllerPaslon) GetPaslon(c echo.Context) error {
	paslon, err := cr.PaslonRepository.GetPaslon()
	if err != nil {
		return c.JSON(http.StatusOK, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get paslon Success", paslon))
}

func (cr *controllerPaslon) GetPaslonById(c echo.Context) error {
	id, _ := strconv.Atoi("id")
	paslon, err := cr.PaslonRepository.GetPaslonById(id)
	if err != nil {
		return c.JSON(http.StatusOK, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get paslon Success", paslon))
}