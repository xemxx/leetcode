package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	m := map[int]int{}

	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	str := string(bytes)
	start := 0
	for i := range str {
		if str[i] == ' ' {
			v, _ := strconv.Atoi(str[start:i])
			if _, ok := m[v]; !ok {
				m[v] = 0
			}
			m[v]++
			start = i + 1
		}
	}
	v, _ := strconv.Atoi(str[start:])
	if _, ok := m[v]; !ok {
		m[v] = 0
	}
	m[v]++

	//fmt.Println(m)
	for k, v := range m {
		if v%2 != 0 {
			fmt.Println(k)
			return
		}
	}
}
