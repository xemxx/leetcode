package main

import "fmt"

func main() {
	solve()
}

func solve() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	max, iMax := arr[0], arr[0]
	for _, v := range arr[1:] {
		tmp := v + iMax
		if tmp < v {
			tmp = v
		}
		iMax = tmp
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(max)
}
