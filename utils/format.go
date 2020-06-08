package utils

import (
	"bytes"
	"encoding/json"
)

func JsonFormat(v interface{}) string {
	bytes, _ := json.MarshalIndent(v, "  ", "  ")
	return string(bytes)
}
func JsonBeautFormat(res []byte) string {
	var buf bytes.Buffer
	_ = json.Indent(&buf, res, " ", "  ")
	return buf.String()
}
