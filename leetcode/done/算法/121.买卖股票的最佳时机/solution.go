/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i]
		if tmp < min {
			min = tmp
			continue
		}
		if max < tmp-min {
			max = tmp - min
		}
	}
	return max
}

// @lc code=end

