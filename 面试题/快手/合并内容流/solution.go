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

	reader = bufio.NewReader(os.Stdin)
	bytes, _, _ = reader.ReadLine()
	str = string(bytes)
	char := []string{}
	start = 0
	for i := range str {
		if str[i] == ' ' {
			char = append(char, str[start:i])
			start = i + 1
		}
	}
	char = append(char, str[start:])

	for i := range now {
		if i%4 == 0 {
			fmt.Print("%v ", char[i])
		} else {
			fmt.Print("%v ", now[i])
		}
	}

}
