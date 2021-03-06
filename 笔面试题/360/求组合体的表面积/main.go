package main

/*
题目： 用一个二维数组表示一个由多个体积为1的正方体组成的组合体（多面体），每个元素的数值n表明此位置存在n个正方体堆叠，请求出其表面积，
输入：2 2
	 2 1
	 1 1
输出：20
*/
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
			}
			if arr[i][j] > buttom {
				m += arr[i][j] - buttom
			}
			if arr[i][j] > left {
				m += arr[i][j] - left
			}
			if arr[i][j] > right {
				m += arr[i][j] - right
			}
		}
	}
	return m
}
