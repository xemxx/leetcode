package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&m, &n)
	result := make([][]uint64, n)
	tmp := make([]uint64, m)
	for j := 0; j < m; j++ {
		tmp[j] = 1
	}
	result[0] = tmp
	for i := 1; i < n; i++ {
		tmp := make([]uint64, m)
		tmp[0] = 1
		for j := 1; j < m; j++ {
			tmp[j] = tmp[j-1] + result[i-1][j]

		}
		result[i] = tmp
	}
	fmt.Println(result[n-1][m-1])
}
