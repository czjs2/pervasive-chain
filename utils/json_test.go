package utils

import (
	"fmt"
	"testing"
)

func TestJosn(t *testing.T){
	value := GetJsonValue("{\"event\": \"block\", \"body\": {}}", "event")
	fmt.Println(value)
}
