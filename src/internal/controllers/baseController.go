package controllers

import (
	"encoding/json"
	"net/http"
)

type Controller interface {
	Post()
	Get()
	Put()
	Delete()
}

type BaseController struct{}

func (baseController *BaseController) SendError(w http.ResponseWriter, error string, status int) {
	http.Error(w, error, status)
}

func (baseController *BaseController) SendResponse(w http.ResponseWriter, content any) {
	json.NewEncoder(w).Encode(content)
}
