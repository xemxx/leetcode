package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution([][]int{[]int{2, 2}, []int{2, 1}, []int{1, 1}}))

}

func solution(arr [][]int) int {
	m := 0
	for i := range arr {
		for j := range arr[i] {

			m += 2
			top, buttom, left, right := 0, 0, 0, 0

			if i == 0 {
				top = 0
			} else {
				top = arr[i-1][j]
			}
			if j == 0 {
				left = 0
			} else {
				left = arr[i][j-1]
			}
			if i == len(arr)-1 {
				buttom = 0
			} else {
				buttom = arr[i+1][j]
			}
			if j == len(arr[i])-1 {
				right = 0
			} else {
				right = arr[i][j+1]
			}

			if arr[i][j] > top {
				m += arr[i][j] - top
			} else if arr[i][j] > buttom {
				m += arr[i][j] - buttom
			} else if arr[i][j] > left {
				m += arr[i][j] - left
			} else if arr[i][j] > right {
				m += arr[i][j] - right
			}

		}
	}
	return m
}
