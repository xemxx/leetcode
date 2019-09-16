package main

import "fmt"

func main() {
	//输入
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	count := 0
	//贪心策略
	for len(arr) > 1 {
		//因为是一圈排列，所以需要考虑首尾合并的情况
		x, y, max := 0, len(arr)-1, arr[0]+arr[len(arr)-1]
		for i := 1; i < len(arr); i++ {
			if arr[i]+arr[i-1] > max {
				max = arr[i] + arr[i-1]
				x, y = i-1, i
			}
		}
		arr[x] = max
		if x == y-1 {
			copy(arr[x:], arr[y:])
		}
		arr = arr[:len(arr)-1]
		count += max
	}
	fmt.Println(count)
}
