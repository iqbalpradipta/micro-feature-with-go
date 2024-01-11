package controller

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/micro-feature-with-go/helpers"
	service "github.com/iqbalpradipta/micro-feature-with-go/services"
	"github.com/labstack/echo/v4"
)

type controllerPartai struct {
	PartaiRepository service.PartaiService
}

func PartaiController(partaiService service.PartaiService) *controllerPartai {
	return &controllerPartai{partaiService}
}

func (cr *controllerPartai) GetPartai(c echo.Context) error {
	partai, err := cr.PartaiRepository.GetPartai()
	if err != nil {
		return c.JSON(http.StatusOK, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Partai Success", partai))
}

func (cr *controllerPartai) GetPartaiById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	partai, err := cr.PartaiRepository.GetPartaiById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Partai By Id Success", partai))
}