package utils

type BaseResult struct {
	Code    int         `json:"code"`              //返回码
	Message string      `json:"message,omitempty"` //错误信息
	Result  interface{} `json:"result"`            //结构体
}
