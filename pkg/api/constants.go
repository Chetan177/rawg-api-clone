package api

import (
	"net/http"
	"rawg/pkg/model"
)

const (
	StatusSuccess = iota + 1
	ValidationError
)

var StatusMap = map[int]model.Response{
	StatusSuccess:   {Msg: "success", Status: http.StatusOK},
	ValidationError: {Msg: "request validation error", Status: http.StatusBadRequest},
}
