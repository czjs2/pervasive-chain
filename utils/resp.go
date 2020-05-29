package utils

import (
	"net/http"
	"pervasive-chain/model"
)

func FailResponse(msg string) *model.Response {
	return &model.Response{
		Code:    http.StatusOK,
		Message: msg,
		Data:    nil,
	}
}

func SuccessResponse(data interface{}) *model.Response {
	return &model.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success",
	}
}
