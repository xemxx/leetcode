package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	count := []int{0, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr[i] = tmp
		for j := i; j > 0; j-- {
			if count[arr[j]] < count[arr[j-1]] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else if count[arr[j]] == count[arr[j-1]] && arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	// result := ""
	// i := 0
	// for i < m && n > 0 {
	// 	if count[arr[i]]*2 <= n {
	// 		n -= count[arr[i]] * 2
	// 		result = fmt.Sprintf("%s%d%d", result, arr[i], arr[i])
	// 	} else {
	// 		i++
	// 	}
	// }
	// for i := 0; i < m-1; i++ {
	// 	if n >= count[arr[i]] && n <= count[arr[i+1]] {
	// 		n -= count[arr[i]]
	// 		result = fmt.Sprintf("%s%d", result, arr[i])
	// 		break
	// 	}
	// }
	result := ""
	i := 0
	for n > 0 && i < m {
		flag := count[arr[i]]
		if n >= flag {
			n -= flag
			result = fmt.Sprintf("%s%d", result, arr[i])
		} else {
			i++
		}
	}
	fmt.Println(result)
}
