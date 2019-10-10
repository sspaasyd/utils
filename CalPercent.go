package utils

import (
	"github.com/dariubs/percent"
	"strconv"
)

//计算数据的百分比,part为分子，total为分母,num为百分数小数点位数
func CalPercent(part, total float64, num int) string {
	f := percent.PercentOfFloat(part, total)
	percent := strconv.FormatFloat(f, 'f', num, 64)
	s := "%"
	return percent + s
}
