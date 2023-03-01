/*
 * @lc app=leetcode.cn id=2373 lang=golang
 *
 * [2373] 矩阵中的局部最大值
 */

// @lc code=start
func largestLocal(grid [][]int) [][]int {
	result := make([][]int, len(grid)-2)
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			m := 0
			for k := 0; k < 3; k++ {
				for d := 0; d < 3; d++ {
					m = max(grid[i+k][j+d], m)
				}
			}
			if result[i] == nil {
				result[i] = make([]int, len(grid)-2)
			}
			result[i][j] = m
		}
	}
	return result

}

// @lc code=end

