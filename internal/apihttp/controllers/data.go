package controllers

import (
	services "aggriddatatable/internal/sevices"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataInterface interface {
	GetByUrl(c echo.Context) error
	MockData(c echo.Context) error
}

type Data struct {
	DataService services.DataInterface
}

func NewData(
	dataService services.DataInterface,
) *Data {
	return &Data{
		DataService: dataService,
	}
}

func (o *Data) GetByUrl(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	urlID := echoContext.Param("id")
	if urlID == "" {
		return echo.ErrBadRequest
	}

	response, err := o.DataService.GetByUrl(ctx, urlID)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	return echoContext.JSON(http.StatusOK, response)
}

func (o *Data) MockData(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	response, err := o.DataService.MockData(ctx)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	return echoContext.JSON(http.StatusOK, response)
}
