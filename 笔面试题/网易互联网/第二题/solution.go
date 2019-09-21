package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	arr := make([]string, t)
	for i := 0; i < t; i++ {
		arr[i] = solvee()
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}

func solvee() string {
	var n, k int
	fmt.Scanln(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	flag := true
	i := 0
	for i = 0; i < n-1; {
		max := i + 1
		for j := i + 1; j <= i+k && j <= n-1; j++ {
			if arr[j] >= arr[max] {
				max = j
			}
		}
		if arr[max] > arr[i] {
			if flag {
				flag = false
				i = max
			} else {
				return "NO"
			}
		} else {
			i = max
		}
	}
	if i == n-1 {
		return "YES"
	}
	return "NO"
}
