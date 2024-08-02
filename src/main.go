package main

import (
	"github.com/digidaniel-tech/go-restapi/internal/api"
	"github.com/digidaniel-tech/go-restapi/internal/controllers"
)

func main() {
	port := ":8080"

	router := api.CreateRouter()

	controller := &controllers.ProductsController{}
	controller.SetupRoute(router)

	router.GetEngine().Run(port)
}
