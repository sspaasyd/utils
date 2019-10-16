package utils

import (
	"log"
	"strconv"
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

//根据开始时间和到期时间,计算下一次的到期时间
func GetNextTime(startTime, expireTime string) string {
	//计算两个时间的时间戳
	parse, _ := time.Parse("2006-01-02 15:04:05", startTime)
	parse2, _ := time.Parse("2006-01-02 15:04:05", expireTime)
	//计算时间差值
	sub := parse2.Sub(parse)

	result := parse2.Add(sub).String()
	return result
}

//获取时间戳毫秒值
func GetTimeStamp() string {
	t := time.Now().UnixNano() / 1e6
	timeStamp := strconv.FormatInt(t, 10)
	return timeStamp
}

//比较时分秒的大小
func LeftIsBigger(left, right string) bool {
	if "" == left || "" == right {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		l := left[i]
		r := right[i]
		if l > r {
			return true
		} else if l < r {
			return false
		}
	}
	return false
}
