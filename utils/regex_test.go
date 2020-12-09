package utils

import (
	"fmt"
	"testing"
)

func TestRegexTest(t *testing.T){
	ok := IsRFC339Time("2012-11-01T22:08:41+00:00")
	fmt.Println(ok)
}
