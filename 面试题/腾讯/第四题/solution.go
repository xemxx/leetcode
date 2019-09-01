package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	max := 0
	//mmp := make(map[int]bool)
	for i := 0; i < n; i++ {

		a, b := i-1, i+1
		tmp := arr[i]
		for a >= 0 && arr[a] >= arr[i] {
			tmp += arr[a]
			a--
		}
		for b < n && arr[b] > arr[i] {
			tmp += arr[b]
			b++
		}
		if b < 0 && arr[b] == arr[i] {
			continue
		}
		tmp *= arr[i]
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(max)
}
