package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(arr[i])
	}
	prev := arr[0]
	result := 0
	for i := 1; i < n; i++ {
		if arr[i] != prev+1 {
			result = i
			break
		}
		prev = arr[i]
	}
	fmt.Println(result)
}
