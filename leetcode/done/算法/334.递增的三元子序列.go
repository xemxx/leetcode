/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */
package leetcode

import "math"

// @lc code=start
func increasingTriplet(nums []int) bool {
	// i, j, k := 0, 1, 2
	if len(nums) < 3 {
		return false
	}
	first, last := nums[0], math.MaxInt
	for _, v := range nums[1:] {
		if v > last {
			return true
		}
		if v > first {
			last = v
		} else {
			first = v
		}
	}
	return false
}

// @lc code=end
