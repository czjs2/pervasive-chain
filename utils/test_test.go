package utils

import (
	"fmt"
	"pervasive-chain/model"
	"reflect"
	"testing"
)

type User struct {
	Name string
}

func TestChain(t *testing.T) {

}

func TestReflet(t *testing.T) {
	cmd := model.Block{}
	DoFiledAndMethod(cmd)

}

func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	param := make(map[interface{}]interface{})
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		//fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
		param[field.Name] = value
	}
	fmt.Println(param)
	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
