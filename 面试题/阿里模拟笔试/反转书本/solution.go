package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&arr[i])
	}
	min := 0
	for i := 1; i < n; i++ {
		if arr[i] != i {
			min += dp(arr, i)
		}
	}
	fmt.Println(min)
}
func dp(arr []int, now int) int {
	x := 0
	for i, v := range arr {
		if v == now {
			x = i
			break
		}
	}
	if arr[len(arr)-1] == arr[x] {
		i, j := now, len(arr)-1
		for i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
		return 1
	}
	i, j := x, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	i, j = now, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return 2
}
