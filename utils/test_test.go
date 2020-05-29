package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string
}
func TestReflet(t *testing.T) {
	of := reflect.TypeOf(User{})
	for i:=0;i<of.NumField();i++{
		fmt.Println(of)
	}
}
