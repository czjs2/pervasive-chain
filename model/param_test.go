package model

import (
	"reflect"
	"testing"
	"time"
)

func TestParam(t *testing.T){
	params := Param{"time":time.Now(),"info":"111"}

}


func NewObj(obj interface{}) interface{} {
	getType := reflect.TypeOf(obj)
	getValue := reflect.ValueOf(obj)
	param := make(map[string]interface{})
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		param[field.Name] = value
	}
	return param
}

