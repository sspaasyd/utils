package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//响应前端数据
func ResponseData(code int, message string, data interface{}, w http.ResponseWriter) {
	br := &BaseResult{
		Code:    code,
		Message: message,
		Result:  data,
	}
	//将数据转为json
	result, e := json.Marshal(br)
	if e != nil {
		fmt.Println("ResponseError将数据转为json异常", e)
		return
	}
	Response(w, result)
}

//设置响应
func Response(response http.ResponseWriter, data []byte) {
	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Request-Method", "*")
	response.Header().Add("Content-Type", "application/json")
	response.Header().Add("Access-Control-Allow-Headers", "User-Agent,X-Requested-With,Cache-Control,Content-Type,"+
		"Access-Token,Authorization,range-token")
	response.Write(data)
}
