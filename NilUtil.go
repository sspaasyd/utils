package utils

import (
	"reflect"
)

//如果结构体为空，返回false,不为空，返回true
func JudgeNil(s interface{}, s2 interface{}) bool {
	b := s == s2
	b2 := reflect.DeepEqual(s, s2)
	if b && b2 {
		//说明结构体为空
		return false
	} else {
		return true
	}

}
