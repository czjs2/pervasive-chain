package service

import "time"

// 毫秒转换为 时间
func millisecondToTime(t int64) time.Time {
	return time.Unix(0, t*1000*1000)
}
