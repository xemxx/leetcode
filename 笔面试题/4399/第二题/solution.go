package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	i, j := 0, 1
	last := false
	for j <= len(s)-1 {
		if s[i] == s[j] {
			last = true
			i++
		} else {
			last = false
			i = 0
		}
		j++
	}
	if last {
		s = s + s[i:]
	} else {
		s = s + s
	}
	fmt.Println(s)
}
