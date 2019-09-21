package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	arr := make([]int, n)
	k := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanln(&k[i])
	}
	res := map[int]int{}
	for _, v := range arr {
		if _, ok := res[v]; ok {
			res[v]++
		} else {
			res[v] = 1
		}
	}
	for _, v := range k {
		fmt.Println(res[v])
	}

}
