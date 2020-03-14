package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	arr := make([]int, k+1)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr[tmp]++
	}
	count := 0
	for i := 1; i <= k; i++ {
		count += ceil(arr[i])
	}
	fmt.Println(count)
}

func ceil(a int) int {
	tmp := a / 2
	if tmp*2 < a {
		return tmp + 1
	}
	return tmp
}
