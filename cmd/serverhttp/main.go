package main

import (
	"aggriddatatable/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {
	c := config.GetConfig()

	router := wireDependencies()

	e := echo.New()
	router.Route(e)

	e.Logger.Fatal(e.Start(":" + c.HTTP_PORT))
}
