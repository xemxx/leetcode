package main

import (
	"fmt"
	"time"
)

func main() {
	var now string
	fmt.Scanln(&now)
	var k int
	fmt.Scanln(&k)
	allT := make([]time.Time, 0)
	for i := 0; i < k; i++ {
		var str string
		fmt.Scanln(&str)
		start := 0
		for i := range str {
			if str[i] == ',' || str[i] == '~' {
				tmp, _ := time.ParseInLocation("2006-01-02T15:04:05", str[start:i], time.Local)
				allT = append(allT, tmp)
				start = i + 1
			}
		}
		tmp, _ := time.ParseInLocation("2006-01-02T15:04:05", str[start:], time.Local)
		allT = append(allT, tmp)
		start = i + 1
	}
	fmt.Println(allT)
}
