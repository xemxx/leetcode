package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	arr := make([]int, t)
	for i := 0; i < t; i++ {
		arr[i] = solvee()
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}

func solvee() int {
	var a, b, q, p int
	fmt.Scanln(&a, &b, &p, &q)
	count := 0
	for a < b {
		if a+p >= b {
			a = a + p
			count++
		} else {
			p *= q
			count++
		}
	}
	return count
}
