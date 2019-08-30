package utils

const (
	SUCCESS         = 200
	SUCCESS_MESSAGE = "操作成功"

	ERROR         = 500
	ERROR_MESSAGE = "服务器异常"

	//注册登录
	REGISTER_SUCCESS         = 1000
	REGISTER_SUCCESS_MESSAGE = "注册成功"

	REGISTER_FAIL         = 1001
	REGISTER_FAIL_MESSAGE = "注册失败"

	NAME_NOT_NIL         = 1002
	NAME_NOT_NIL_MESSAGE = "用户名不能为空"

	EMAIL_NOT_NIL         = 1003
	EMAIL_NOT_NIL_MESSAGE = "邮箱不能为空"

	PASSWORD_NOT_NIL         = 1004
	PASSWORD_NOT_NIL_MESSAGE = "密码不能为空"

	CONFIRMPASSWORD_NOT_NIL         = 1005
	CONFIRMPASSWORD_NOT_NIL_MESSAGE = "确认密码不能为空"

	PASSWORD_NOT_SAME         = 1006
	PASSWORD_NOT_SAME_MESSAGE = "密码不相同"

	PHONE_NOT_NIL         = 1007
	PHONE_NOT_NIL_MESSAGE = "手机号不能为空"

	PHONECODE_NOT_NIL         = 1008
	PHONECODE_NOT_NIL_MESSAGE = "手机验证码不能为空"
)
