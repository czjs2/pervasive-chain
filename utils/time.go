package utils

import (
	"fmt"
	"strconv"
	"time"
)

// 时间格式化参数
const SysTimefrom = "2006-01-02 15:04:05"

func GetTime() time.Time {
	return time.Now()
}

func GetNowTime() string {
	now := time.Now()
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}
	hour := strconv.Itoa(now.Hour())
	if len(hour) == 1 {
		hour = "0" + hour
	}
	minute := strconv.Itoa(now.Minute())
	if len(minute) == 1 {
		minute = "0" + minute
	}
	second := strconv.Itoa(now.Second())
	if len(second) == 1 {
		second = "0" + second
	}
	return fmt.Sprintf("%d-%s-%s %s:%s:%s", now.Year(), month, day, hour, minute, second)
}

func GetCurZeroTime() (time.Time, error) {
	today := GetCurrentDay()
	return ParseLocalTime(today)
}

func ParseLocalTime(t string) (time.Time, error) {
	//中国时区
	location, err := time.LoadLocation("Asia/Shanghai")
	ntime, err := time.ParseInLocation(SysTimefrom, t, location)
	return ntime, err
}

// 获取当天字符串
func GetCurrentDay() string {
	now := time.Now()
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}
	return fmt.Sprintf("%d-%s-%s 00:00:00", now.Year(), month, day)
}
