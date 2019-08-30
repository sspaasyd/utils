package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonFile(path string, data interface{}) interface{} {
	fmt.Println("path:", path)
	file, e := os.Open(path)
	if e != nil {
		fmt.Println("ReadJsonFile open json file error:" + e.Error())
		return data
	}
	decoder := json.NewDecoder(file)
	e2 := decoder.Decode(data)
	if e2 != nil {
		fmt.Println("ReadJsonFile json decode error:" + e2.Error())
		return data
	}
	return data
}
