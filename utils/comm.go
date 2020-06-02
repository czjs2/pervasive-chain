package utils

import (
	"bytes"
	"fmt"
)

func MapToStr(m map[string]interface{}) string {
	var buf bytes.Buffer
	for k, v := range m {
		buf.WriteString("[")
		buf.WriteString(k)
		buf.WriteString(":")
		buf.WriteString(fmt.Sprintf("%v", v))
		buf.WriteString("]")
	}
	return buf.String()
}
