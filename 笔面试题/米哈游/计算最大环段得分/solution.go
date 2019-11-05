package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 1; i < n; i++ {
		arr = append(arr, arr[i])
	}
	n = len(arr)
	f := make([]int, n)
	max := arr[0]
	f[0] = arr[0]
	for i := 1; i < n; i++ {
		tmpmax := arr[i] + f[i-1]
		if arr[i] > tmpmax {
			tmpmax = arr[i]
		}
		f[i] = tmpmax
		if tmpmax > max {
			max = tmpmax
		}
	}
	fmt.Println(max)
}
