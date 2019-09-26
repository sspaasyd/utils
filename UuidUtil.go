package utils

import (
	. "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

//获取32位uuid随机数
func Get32UUID() string {
	uuids := NewV4()
	id := strings.Replace(uuids.String(), "-", "", -1)
	return id
}

const (
	KC_RAND_KIND_NUM   = 0  // 纯数字
	KC_RAND_KIND_LOWER = 1  // 小写字母
	KC_RAND_KIND_UPPER = 2  // 大写字母
	KC_RAND_KIND_ALL   = 3  // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}

//获取指定长度的随机字符串
func GetRandomString(i int)string{
	str:="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	char:=[]byte(str)
	var result string
	rand.Seed(time.Now().UnixNano())
	for j:=0;j<i;j++{
		intn := rand.Intn(len(str) - 1)
		s := string(char[intn])
		result = result + s
	}
	return result
}