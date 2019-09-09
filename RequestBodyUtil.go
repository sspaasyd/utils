package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestBody(r *http.Request, v *interface{}) error {
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		fmt.Println("RequestBody,ioutil.ReadAll(r.Body)异常", e)
		return e
	}

	//关闭
	defer r.Body.Close()

	//如果bytes为空
	if data == nil {
		fmt.Println("RequestBody,data为空")
		return nil
	}

	//解析为结构体
	e2 := json.Unmarshal(data, v)
	if e2 != nil {
		fmt.Println("RequestBody,解析为结构体异常", e2)
		return e2
	}
	return nil
}
