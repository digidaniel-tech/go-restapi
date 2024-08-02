package api

import "github.com/gin-gonic/gin"

type Router struct {
	engine *gin.Engine
}

func CreateRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *Router) SetupRoute(method, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		r.engine.GET(path, handler)
	case "POST":
		r.engine.POST(path, handler)
	case "PUT":
		r.engine.PUT(path, handler)
	case "DELETE":
		r.engine.DELETE(path, handler)
	}
}
