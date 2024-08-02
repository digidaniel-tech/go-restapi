package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Post()
	Get()
	Put()
	Delete()
}

type BaseController struct{}

func (baseController *BaseController) SendResponse(context *gin.Context, status int, content any) {
	context.JSON(status, content)
}
