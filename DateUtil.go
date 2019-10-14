package utils

import (
	"log"
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

//判断指定时间是否小于当前时间，是返回true.否则返回false
func CompareTime(t string) bool {
	now := GetNow()
	now2, e1 := time.Parse("2006-01-02 15:04:05", now)
	t2, e2 := time.Parse("2006-01-02 15:04:05", t)

	if e1 != nil || e2 != nil {
		log.Println("e1", e1)
		log.Println("e2", e2)
	}
	b := t2.Before(now2)
	return b
}
