package main

import (
	"aggriddatatable/internal/apihttp"
	"aggriddatatable/internal/apihttp/controllers"
	services "aggriddatatable/internal/sevices"
)

func wireDependencies() *apihttp.Router {
	// services
	dataService := services.NewData()

	// controllers
	dataController := controllers.NewData(dataService)

	return apihttp.NewRouter(
		dataController,
	)
}
