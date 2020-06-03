package utils

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
}
func TestReflet(t *testing.T) {
	rand := Rand(1000)
	fmt.Println(rand)
}
