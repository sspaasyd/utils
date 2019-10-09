package utils

import (
	"reflect"
)

//把一个struct中的属性复制到另一个struct中
//源数source据结构可以用值传参，目标target结构体用引用传参,例：&target
func CopyProperties(source interface{}, target interface{}) {
	getSourceValue := reflect.ValueOf(source)        //源数据结构
	getSourceType := reflect.TypeOf(source)          //源数据结构类型
	getTargetValue := reflect.ValueOf(target).Elem() //目标数据结构
	for i := 0; i < getTargetValue.NumField(); i++ {
		targetField := getTargetValue.Type().Field(i)
		targetFieldValue := getTargetValue.FieldByName(targetField.Name)
		if _, ok := getSourceType.FieldByName(targetField.Name); ok {
			targetFieldValue.Set(getSourceValue.FieldByName(targetField.Name))
		}
	}
}
