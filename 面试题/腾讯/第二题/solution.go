package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		arr[i] = append(arr[i], a, b)
	}
	count := 0
	quickSort(arr, 0, n-1)
	for i := 0; i < n; i++ {
		count += arr[i][0]*(i+1-1) + arr[i][1]*(n-(i+1))
	}
	fmt.Println(count)
}

func quickSort(arr [][]int, left, right int) {
	if left < right {
		pivot := paritition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}
func paritition(arr [][]int, left, right int) int {
	temp, j := left, left+1
	for i := j; i <= right; i++ {
		if arr[i][1] >= arr[temp][1] && arr[i][0] <= arr[temp][0] {
			swap(arr, i, j)
			j++
		}
	}
	swap(arr, left, j-1)
	return j - 1
}

func swap(arr [][]int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
