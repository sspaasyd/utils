package utils

import (
	"fmt"
	"time"
)

func GetNow() string {
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")
	fmt.Println(date)
	return date
}
