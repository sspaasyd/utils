package utils

import (
	"fmt"
	"strconv"
)

//小数点位数处理

//先通过Sprintf保留两位小数，再转成float64
func Decimal(f float64, num int) float64 {
	s1 := "%."
	s2 := "f"
	itoa := strconv.Itoa(num)
	s := s1 + itoa + s2
	result, _ := strconv.ParseFloat(fmt.Sprintf(s, f), 64)
	return result
}
