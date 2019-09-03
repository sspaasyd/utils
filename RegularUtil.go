package utils

import "regexp"

// 手机正则
var phone_el = "^1[3|4|5|7|8]\\d{9}$"

//用户名正则
var name_el = "^[a-zA-Z][a-z0-9]{3,14}$"

//密码正则
var password_el = "^[a-zA-Z][a-zA-Z0-9_.+]{7,19}$"

//邮箱正则
var email_el = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)$"

//验证手机号是否合法
func RegularPhone(phone string) bool {
	compile := regexp.MustCompile(phone_el)
	b := compile.MatchString(phone)
	return b
}

//验证用户名是否合法
func RegularName(name string) bool {
	compile := regexp.MustCompile(name_el)
	b := compile.MatchString(name)
	return b
}

//验证密码是否合法
func RegularPassword(password string) bool {
	compile := regexp.MustCompile(password_el)
	b := compile.MatchString(password)
	return b
}

//验证邮箱是否合法
func RegularEmail(email string) bool {
	compile := regexp.MustCompile(email_el)
	b := compile.MatchString(email)
	return b
}
