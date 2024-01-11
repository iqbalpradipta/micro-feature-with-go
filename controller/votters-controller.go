package controller

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/micro-feature-with-go/helpers"
	service "github.com/iqbalpradipta/micro-feature-with-go/services"
	"github.com/labstack/echo/v4"
)

type controllerVotters struct {
	VottersRepository service.VottersService
}

func VottersController(vottersService service.VottersService) *controllerVotters {
	return &controllerVotters{vottersService}
}

func (cr *controllerVotters) GetVotters(c echo.Context) error {
	votters, err := cr.VottersRepository.GetVotters()
	if err != nil {
		return c.JSON(http.StatusOK, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Votters Success", votters))
}

func (cr *controllerVotters) GetVottersById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	votters, err := cr.VottersRepository.GetVottersById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Votters By Id Success", votters))
}