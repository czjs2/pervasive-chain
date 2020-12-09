package model

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	imap := Param{}
	imap["shijian"] = time.Now()
	imap["height"] = 11
	var input interface{}
	input = imap
	m := reflect.ValueOf(input)
	if m.Kind() == reflect.Map {
		res := reflect.MakeMap(m.Type())
		keys := m.MapKeys()
		for _, k := range keys {
			key := k.Convert(res.Type().Key()) //.Convert(m.Type().Key())
			value := m.MapIndex(key)
			res.SetMapIndex(key, value)
			fmt.Println(key,value)
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
