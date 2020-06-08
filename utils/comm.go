package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
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

func Rand(src int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(src)+1
}
