package utils

import (
	"time"
)

//获取当前时间
func GetNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//获取指定时间(年,月,日)前后的时间
func GetTime(year, month, day int) string {
	return time.Now().AddDate(year, month, day).Format("2006-01-02 15:04:05")

}

//获取下一个小时的时间
func GetNextHourTime() string {
	return time.Now().Add(-time.Hour).Format("2006-01-02 15:04:05")
}
