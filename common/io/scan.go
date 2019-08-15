package io

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// ScanTwoArr 用于获取某些oj系统的二维数组的输入.
func ScanTwoArr() [][]int {
	arr := [][]int{}
	for {
		reader := bufio.NewReader(os.Stdin)
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(bytes)
		now := []int{}
		start := 0
		for i := range str {
			if str[i] == ' ' {
				num, _ := strconv.Atoi(str[start:i])
				now = append(now, num)
				start = i + 1
			}
		}
		num, _ := strconv.Atoi(str[start:])
		now = append(now, num)
		arr = append(arr, now)
	}
	return arr
}
