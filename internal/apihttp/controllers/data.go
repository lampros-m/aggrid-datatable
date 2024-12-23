package controllers

import (
	services "aggriddatatable/internal/sevices"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataInterface interface {
	GetAll(c echo.Context) error
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

func (o *Data) GetAll(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	response, err := o.DataService.GetAll(ctx)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	return echoContext.JSON(http.StatusOK, response)
}
