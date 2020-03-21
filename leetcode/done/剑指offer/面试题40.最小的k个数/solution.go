package main

import "fmt"

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}
func getLeastNumbers(arr []int, k int) []int {
	result := arr[:k]
	for i := k/2 - 1; i >= 0; i-- {
		shiftDown(result, i, k-1)
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < result[0] {
			result[0] = arr[i]
			shiftDown(result, 0, k-1)
		}
	}
	return result
}

func shiftDown(arr []int, i, l int) {
	for i*2+1 <= l {
		left, right := i*2+1, i*2+2
		if right <= l {
			if arr[right] > arr[left] {
				left = right
			}
		}
		if arr[i] < arr[left] {
			arr[i], arr[left] = arr[left], arr[i]
			i = left
		} else {
			break
		}
	}
}
