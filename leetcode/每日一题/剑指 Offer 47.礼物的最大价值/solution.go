// @algorithm @lc id=100327 lang=golang
// @title li-wu-de-zui-da-jie-zhi-lcof

package main

func maxValue(grid [][]int) int {
	result := make([][]int, len(grid))
	for i, v := range grid {
		result[i] = make([]int, len(v))
		result[i][0] = v[0]
		if i > 0 {
			result[i][0] += result[i-1][0]
		}
		for j := 1; j < len(v); j++ {
			if i == 0 {
				result[i][j] = result[i][j-1] + v[j]
			} else {
				result[i][j] = max(result[i][j-1], result[i-1][j]) + v[j]
			}
		}
	}

	return result[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, x, y int) int {
	// 回溯，超时，直接dfs也超时
	if x == 0 && y == 0 {
		return grid[x][y]
	} else if x == 0 {
		return dfs(grid, x, y-1) + grid[x][y]
	} else if y == 0 {
		return dfs(grid, x-1, y) + grid[x][y]
	} else {
		return max(dfs(grid, x-1, y), dfs(grid, x, y-1)) + grid[x][y]
	}
}
