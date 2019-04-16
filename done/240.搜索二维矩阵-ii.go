package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 5))

}
func searchMatrix(matrix [][]int, target int) bool {
	res := false
	if len(matrix) > 0 {
		for _, t := range matrix {
			if len(t) > 0 {
				if target >= t[0] {
					for _, v := range t {
						if target < v {
							break
						}
						if v == target {
							res = true
						}
					}
				}
			}
		}
	}
	return res
}

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (36.51%)
 * Total Accepted:    12.8K
 * Total Submissions: 35.2K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n5'
 *
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
 *
 *
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 *
 *
 * 示例:
 *
 * 现有矩阵 matrix 如下：
 *
 * [
 * ⁠ [1,   4,  7, 11, 15],
 * ⁠ [2,   5,  8, 12, 19],
 * ⁠ [3,   6,  9, 16, 22],
 * ⁠ [10, 13, 14, 17, 24],
 * ⁠ [18, 21, 23, 26, 30]
 * ]
 *
 *
 * 给定 target = 5，返回 true。
 *
 * 给定 target = 20，返回 false。
 *
 */
