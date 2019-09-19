package utils

import (
	"fmt"
	"time"
)

//获取当前时间
func GetNow() string {
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")
	fmt.Println(date)
	return date
}

//获取指定时间(年,月,日)前后的时间
func GetTime(year, month, day int) string {
	date := time.Now().AddDate(year, month, day).Format("2006-01-02 15:04:05")
	return date
}
