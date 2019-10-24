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
	return time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")
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

//获取两个日期之间的天数，天数取整
//date 日期
func GetDaysTwoTime(startTime string, endTime string) (days int) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	//把startTime,endTime转换成当前时区的时间
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", startTime, location)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, location)
	t3 := t2.Unix() - t1.Unix()
	//获取整数天数
	strDays := strconv.FormatInt(t3/(3600*24), 10)
	days, _ = strconv.Atoi(strDays)
	return
}

//获取两个日期之间的天数，如果天数不为整那天数加1
func DayBetweenTwoTime(startTime, endTime string) int {
	location, _ := time.LoadLocation("Asia/Shanghai")
	//把startTime,endTime转换成当前时区的时间
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", startTime, location)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, location)
	t3 := t2.Unix() - t1.Unix()
	var fdays int64
	//开始时间与结束时间之间的秒数差值是否是整天数，如果不是整天数，最后天数加1
	if t3%(3600*24) == 0 {
		fdays = t3 / (3600 * 24)
	} else {
		fdays = (t3 / (3600 * 24)) + 1
	}

	days, _ := strconv.Atoi(strconv.FormatInt(fdays, 10))
	return days
}

//结束时间与当前时间相差的天数
func DayBetweenTwoTimeEnd(endTime string) int {
	location, _ := time.LoadLocation("Asia/Shanghai")
	//把endTime转换成当前时区的时间
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, location)
	t2 := t1.Unix() - time.Now().Unix()
	var fdays int64
	//结束时间与当前时间之间的秒数差值是否是整天数，如果不是整天数，最后天数加1
	if t2%(3600*24) == 0 {
		fdays = t2 / (3600 * 24)
	} else {
		fdays = (t2 / (3600 * 24)) + 1
	}
	days, _ := strconv.Atoi(strconv.FormatInt(fdays, 10))
	return days
}
