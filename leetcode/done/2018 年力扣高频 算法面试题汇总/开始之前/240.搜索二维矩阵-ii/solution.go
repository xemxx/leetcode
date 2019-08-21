package main

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for _, t := range matrix {
		for _, v := range t {
			if target < v {
				break
			}
			if v == target {
				return true
			}
		}
	}
	return false
}
