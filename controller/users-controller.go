package controller

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/micro-feature-with-go/helpers"
	service "github.com/iqbalpradipta/micro-feature-with-go/services"
	"github.com/labstack/echo/v4"
)

type controllerUsers struct {
	UsersRepository service.UsersService
}

func UsersController(usersService service.UsersService) *controllerUsers {
	return &controllerUsers{usersService}
}

func (cr *controllerUsers) GetUsers(c echo.Context) error {
	users, err := cr.UsersRepository.GetUsers()
	if err != nil {
		return c.JSON(http.StatusOK, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Users Success", users))
}

func (cr *controllerUsers) GetUsersById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := cr.UsersRepository.GetUsersById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Users By Id Success", users))
}