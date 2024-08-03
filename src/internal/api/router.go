package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func CreateRouter() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}

func (r *Router) GetEngine() *mux.Router {
	return r.router
}

func (r *Router) SetupRoute(method, path string, handler http.Handler) {
	switch method {
	case "GET":
		r.router.Handle(path, handler).Methods("GET")
	case "POST":
		r.router.Handle(path, handler).Methods("POST")
	case "PUT":
		r.router.Handle(path, handler).Methods("PUT")
	case "DELETE":
		r.router.Handle(path, handler).Methods("DELETE")
	}
}
