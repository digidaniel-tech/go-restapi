package controllers

import (
	"github.com/digidaniel-tech/go-restapi/internal/api"
	"github.com/digidaniel-tech/go-restapi/internal/usecases"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	BaseController
}

func (controller *ProductsController) SetupRoute(router *api.Router) {
	router.SetupRoute("GET", "/products", controller.GetProducts)
	router.SetupRoute("GET", "/products/:id", controller.GetProduct)
}

func (controller *ProductsController) GetProduct(context *gin.Context) {
	id := context.Param("id")
	controller.SendResponse(context, 200, usecases.GetProduct(id))
}

func (controller *ProductsController) GetProducts(context *gin.Context) {
	controller.SendResponse(context, 200, usecases.GetProducts())
}
