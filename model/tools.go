package model

import (
	"reflect"
	"time"
)

func ObjToMap(obj interface{}) interface{}{
	getType := reflect.TypeOf(obj)
	getValue := reflect.ValueOf(obj)
	param := make(map[string]interface{})
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		if field.Type.String()=="time.Time"{
			value := getValue.Field(i).Interface()
			param[field.Name] = value.(time.Time).Local()
		}
	}
	return param
}
