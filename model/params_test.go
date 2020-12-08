package model

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	param := Param{}
	param["time"] = time.Now()
	param["height"] = 11
	for k,v :=range param{
		typeOf := reflect.TypeOf(v)
		fmt.Println(typeOf.Name(),k)
	}
}

func send(param Param) {
	m := reflect.ValueOf(param)
	if m.Kind() == reflect.Map {
		res := reflect.MakeMap(m.Type())
		keys := m.MapKeys()
		for _, k := range keys {
			fmt.Println(k.Type())
			of := reflect.TypeOf(k)
			fmt.Println(of)
			key := k.Convert(res.Type().Key()) //.Convert(m.Type().Key())
			value := m.MapIndex(key)
			fmt.Println(key, value)
			res.SetMapIndex(key, value)
		}
	}
}

func NewObj(obj interface{}) interface{} {
	getType := reflect.TypeOf(obj)
	getValue := reflect.ValueOf(obj)
	param := make(map[string]interface{})
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Println(field,value)
		param[field.Name] = value
	}
	return param
}
