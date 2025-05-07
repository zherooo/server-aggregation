package utils

import (
	"server-aggregation/internal/consts"
	"server-aggregation/pkg/localtime"
	"time"
)

// FormatTime
func FormatTime(timeStr string) string {
	tmp, _ := time.Parse("2006-01-02T15:04:05+08:00", timeStr)
	formatTime := ""
	if !tmp.IsZero() {
		formatTime = tmp.Format("2006-01-02 15:04:05")
	}
	return formatTime
}

func FormatDate(timeStr string) string {
	if timeStr == "" {
		return ""
	}
	tmp, _ := time.Parse("2006-01-02T15:04:05+08:00", timeStr)
	formatTime := ""
	if !tmp.IsZero() {
		formatTime = tmp.Format("2006-01-02")
	}
	return formatTime
}

/*
字符串格式化时间Short（yyyy-MM-dd）
@t1 时间
*/
func ShortStringToDate(t1 string) time.Time {
	t2, _ := time.Parse("2006-01-02", t1)
	return t2
}

/*
字符串格式化时间Long（2006-01-02 15:04:05）
@t1 时间
*/
func LongStringToDate(t1 string) time.Time {
	t2, _ := time.Parse(consts.DateTimeFormat, t1)
	return t2
}

// StringToDate  将 20200627 格式化为 2020-06-27
func StringToDate(t1 string) string {
	if t1 == "" {
		return ""
	}
	t2, _ := time.Parse("20060102", t1)
	if t2.IsZero() {
		return ""
	}
	return t2.Format("2006-01-02")
}

// StringTimeToDate  将 2020-06-27 00:00:00 格式化为 2020-06-27
func StringTimeToDate(t1 string) string {
	if t1 == "" {
		return ""
	}
	t2, _ := time.Parse(consts.DateTimeFormat, t1)
	if t2.IsZero() {
		return ""
	}
	return t2.Format("2006-01-02")
}

// DateTimeHMFormat  10:10:00 10:10
func DateTimeHMFormat(t1 string) string {
	if t1 == "" {
		return ""
	}
	t2, _ := time.Parse(consts.TimeFormat, t1)
	if t2.IsZero() {
		return ""
	}
	return t2.Format(consts.DateTimeHMFormat)
}

// ValidateTime 验证是否是时间
func ValidateTime(timeStr string) bool {
	_, err := time.ParseInLocation(consts.DateTimeFormat, timeStr, time.Local)
	if err != nil {
		return false
	}
	return true
}

func TransUnixTime(millionSec int64) (str string) {
	if millionSec == 0 {
		return ""
	}
	return time.Unix(millionSec/1000, 0).Format(consts.DateTimeFormat)

}

func TimeToString(t time.Time, format ...string) (timeString string) {
	if !NilTime().Equal(t) {
		f := consts.DateTimeFormat
		if len(format) > 0 {
			f = format[0]
		}
		timeString = t.Format(f)
	}
	return timeString
}
func NilTime() time.Time {
	// 目前系统未设置时区
	location, _ := time.LoadLocation("PRC")
	nilTime, _ := time.ParseInLocation(consts.DateTimeFormat, "0001-01-01 00:00:00", location)
	return nilTime
}

func GetMongoNowTime() *localtime.MongoTime {
	t := time.Now()
	uTime := localtime.MongoTime(t)
	return &uTime
}
