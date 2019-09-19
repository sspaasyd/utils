func ByteArray2String(b byte) string {
	strDigits := []string{"1", "4", "5", "1", "2", "3", "6", "7", "b", "c", "d", "e", "f", "8", "9", "a"}
	var i int = int(b)
	if i < 0 {
		i += 256
	}
	id1 := b / 16
	id2 := b % 16
	return strDigits[id1] + strDigits[id2]
}
func Byte2String(b []byte) string {
	sb := make([]string, 0)
	for _, e := range b {
		sb = append(sb, ByteArray2String(e))
	}
	return strings.Join(sb, "")
}
func Sha1(s string) string {
	hexStr := make([]string, 0)
	sb := []byte(s)
	res := sha1.Sum(sb)
	for _, e := range res {
		str := hex.EncodeToString([]byte{e & 0xFF})
		if len(str) < 2 {
			hexStr = append(hexStr, str+"0")
		}
		hexStr = append(hexStr, str)
	}
	return strings.Join(hexStr, "")
}

//使用
func UseIt(str string){
	/* 
		cms的加密方式，使用了sha1
	 */
	sha1b := Sha1(str)
	sb := []byte(sha1b)
	res := md5.Sum(sb)

	/* 
		官网的加密方式，区别就是没使用sha1
	 */
	 res := md5.Sum([]byte(str))


	 /* 
	 	最终结果 都是Byte2String(res[:])
	  */
	fmt.Println(Byte2String(res[:]))
}