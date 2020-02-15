package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(solve(s1, s2))
}

func solve(str1, str2 string) int {
	s1, s2 := getArr(str1), getArr(str2)
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] > s2[i] {
			return 1
		} else if s1[i] < s2[i] {
			return -1
		}
	}
	return 0
}

func getArr(str string) []int {
	arr := []int{}
	start := 0
	for i := range str {
		if str[i] == '.' || str[i] == ',' {
			num, _ := strconv.Atoi(str[start:i])
			arr = append(arr, num)
			start = i + 1
		}
	}
	num, _ := strconv.Atoi(str[start:])
	return append(arr, num)
}
