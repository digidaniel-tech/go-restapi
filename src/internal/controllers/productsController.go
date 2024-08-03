package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/digidaniel-tech/go-restapi/internal/api"
	"github.com/digidaniel-tech/go-restapi/internal/usecases"
)

type ProductsController struct {
	BaseController
}

func (controller *ProductsController) SetupRoute(router *api.Router) {
	router.SetupRoute("GET", "/products", http.HandlerFunc(controller.GetProducts))
	router.SetupRoute("GET", "/products/{id:[0-9]+}", http.HandlerFunc(controller.GetProduct))
}

func (controller *ProductsController) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	product, err := usecases.GetProduct(id)

	// if product == nil {
	// 	controller.SendError(w, fmt.Sprintf("Could not find product with id %s", id), http.StatusNotFound)
	// }

	if err != nil {
		controller.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	controller.SendResponse(w, product)
}

func (controller *ProductsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := usecases.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}
