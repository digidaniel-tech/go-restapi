package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/digidaniel-tech/go-restapi/internal/api"
	"github.com/digidaniel-tech/go-restapi/internal/controllers"
)

var (
	ADDR = "127.0.0.1"
	PORT = "8080"
)

func main() {
	router := api.CreateRouter()

	controller := &controllers.ProductsController{}
	controller.SetupRoute(router)

	url := fmt.Sprintf("%s:%s", ADDR, PORT)
	srv := &http.Server{
		Handler:      router.GetEngine(),
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listen on " + url)
	log.Fatal(srv.ListenAndServe())
}
