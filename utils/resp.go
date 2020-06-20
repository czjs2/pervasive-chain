package utils

import (
	"net/http"
	"pervasive-chain/code"
	"pervasive-chain/model"
)

func FailResponse(msg string) *model.Response {
	return &model.Response{
		Code:    code.StatusFail,
		Message: msg,
		Data:    nil,
	}
}

func ResponseWithCode(code int, msg string, data interface{}) *model.Response {
	return &model.Response{
		Code:    code,
		Data:    data,
		Message: msg,
	}
}

func SuccessResponse(data interface{}) *model.Response {
	return &model.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "OK",
	}
}
