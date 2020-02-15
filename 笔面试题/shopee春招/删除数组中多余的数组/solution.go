package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
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

	solve(now)
}

func solve(now []int) {
	start, end := 0, 0
	l := 0
	//fmt.Println(now)
	for i := 0; i < len(now); i++ {
		if end == now[i] {
			if start < end {
				l++
				start++
			}
		} else {
			end = now[i]
			start = 1
			l++
		}
		// fmt.Println(start, end, l, now[i])
	}
	fmt.Println(l)
}
