package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanln(&k, &n)
	arr := make([][]int, n)
	var root int
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		tmp := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&tmp[j])
		}
		if tmp[0] == 0 {
			root = i
		}
		arr[i] = tmp
	}
	count := 0
	val := []int{root}
	for len(val) > 0 {
		tmp := val[0]
		if len(arr[tmp])-1 >= k {
			count++
			val = arr[tmp][1:]
		} else {
			nodes := arr[tmp][1:]
			sort.Sort(sort.Reverse(sort.IntSlice(nodes)))
			val = append(val, nodes...)
		}

	}
	fmt.Println(count)
}
