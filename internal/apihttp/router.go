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

	e.GET("/dataurl/:id", o.Data.GetByUrl)
	e.GET("/mockdataalpha", o.Data.MockData)
	e.GET("/mockdatabeta", o.Data.MockData)
	e.GET("/mockdatagamma", o.Data.MockData)
	e.GET("/index", func(c echo.Context) error { return c.File("assets/index.html") })
}
