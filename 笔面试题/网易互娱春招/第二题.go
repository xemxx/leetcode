package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	w := make([]int, n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	fmt.Scan(&m)
	sort.Ints(w)
	sort.Ints(t)
	i := 0
	j := 0
	result := 1
	for i < n {
		if t[i] > w[i] {
			fmt.Println(0)
			return
		}
		for j < n && t[j] <= w[i] {
			j++
		}
		result *= (j - i) % m
		j--
		i++
	}
	fmt.Println(result)
}
