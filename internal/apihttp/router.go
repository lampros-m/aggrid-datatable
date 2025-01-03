package apihttp

import (
	"aggriddatatable/internal/apihttp/controllers"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Data controllers.DataInterface
}

func NewRouter(
	data controllers.DataInterface,
) *Router {
	return &Router{
		Data: data,
	}
}

func (o *Router) Route(e *echo.Echo) {
	e.Static("/", "assets")

	e.GET("/data", o.Data.GetAll)
	e.GET("/mockdata", o.Data.MockData)
	e.GET("/index", func(c echo.Context) error { return c.File("assets/index.html") })
}
