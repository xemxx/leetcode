/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */
package leetcode

import "math"

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp_i10, dp_i11, dp_i20, dp_i21 := 0, math.MinInt64, 0, math.MinInt64
	for _, v := range prices {
		dp_i10 = max(dp_i10, dp_i11+v)
		dp_i11 = max(dp_i11, -v)
		dp_i20 = max(dp_i20, dp_i21+v)
		dp_i21 = max(dp_i21, dp_i10-v)

	}
	return dp_i20
}

func max(a, b int) int {
	if b > a {
		a = b
	}
	return a
}

// @lc code=end
