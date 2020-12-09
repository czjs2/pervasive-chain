package utils

import (
	"regexp"
	"time"
)

var (
	//2012-11-01T22:08:41+00:00
	rfc339TimeRegex = regexp.MustCompile(`20[0-9]{2}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\+[0-9]{2}:[0-9]{2}]`)
)

func IsRFC339Time(t string) bool {
	// todo 正则
	_, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return false
	}
	return true
}
