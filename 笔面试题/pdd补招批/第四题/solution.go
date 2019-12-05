package main

import "fmt"

func main() {

	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		//输入
		var m, n int
		fmt.Scan(&n, &m)
		arr := map[int]int{}
		if m != 0 {
			tmp := make(map[int]int, m)
			for j := 0; j < m; j++ {
				var k int
				fmt.Scan(&k)
				tmp[k] = 0
			}
			for _, v := range tmp {
				fmt.Scan(&v)
			}
			arr = tmp
		}
		//计算
		have := 0
		result := 1
		mid := 0
		if len(arr) != 0 {
			// 需要判断前后
			for j := 1; j <= n; j++ {
				tmp := 4
				if v, ok := arr[j]; ok && v != 0 {
					continue
				}
				l := 0
				if v, ok := arr[j-1]; ok && v != 0 {
					tmp--
					l = v
				}
				if v, ok := arr[j-1]; ok && v != 0 {
					if l != 0 && l != v {
						tmp--
					}
				}
				result *= tmp
				result %= 1000000007
				have = 1
			}
		} else {
			for j := 1; j <= n; j++ {
				var tmp int
				if mid == 0 {
					tmp = 4
					mid++
				} else {
					tmp = 3
				}
				have = 1
				result *= tmp
				result %= 1000000007
			}
		}
		if have == 1 {
			fmt.Println(result)
		} else {
			fmt.Println(0)
		}
	}

}
